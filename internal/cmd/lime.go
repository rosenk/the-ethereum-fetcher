package cmd

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ju-popov/the-ethereum-fetcher/internal/bootstrap"
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	"github.com/ju-popov/the-ethereum-fetcher/internal/logbuilder"
	"github.com/ju-popov/the-ethereum-fetcher/internal/shutdownhandler"
	"github.com/ju-popov/the-ethereum-fetcher/internal/validatorbuilder"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
	"github.com/sumup-oss/go-pkgs/task"
)

const appName = "the-ethereum-fetcher"

func NewLimeCommand() *cobra.Command {
	appCommand := &cobra.Command{
		Use:          appName,
		SilenceUsage: true,
	}

	appCommand.PersistentFlags().String("config", "", "config file")

	appCommand.RunE = func(cmd *cobra.Command, args []string) error {
		validate, err := validatorbuilder.Build()
		if err != nil {
			err := errors.Wrap(err, "create config validator")

			fmt.Printf("❌ %s: %v\n", logMessageCommandError, err) //nolint:forbidigo

			return err
		}

		configFilename := cmd.Flags().Lookup("config").Value.String()

		conf, err := config.ReadConfig(configFilename)
		if err != nil {
			err := errors.Wrap(err, "read config")

			fmt.Printf("❌ %s: %v\n", logMessageCommandError, err) //nolint:forbidigo

			return err
		}

		if err := validate.Struct(conf); err != nil {
			err := errors.Wrap(err, "validate config")

			fmt.Printf("❌ %s: %v\n", logMessageCommandError, err) //nolint:forbidigo

			return err
		}

		logConfig := conf.Logger

		log, err := logbuilder.NewBuilder().
			WithLevel(*logConfig.Level).
			WithEncoding(*logConfig.Encoding).
			WithStdoutEnabled(*logConfig.StdoutEnabled).
			WithSyslogEnabled(*logConfig.SyslogEnabled).
			WithSyslogFacility(*logConfig.SyslogFacility).
			WithSyslogTag(*logConfig.SyslogTag).
			Build()
		if err != nil {
			err := errors.Wrap(err, "create logger")

			fmt.Printf("❌ %s: %v\n", logMessageCommandError, err) //nolint:forbidigo

			return err
		}

		return runLimeCommand(log, validate, conf)
	}

	return appCommand
}

//nolint:funlen
func runLimeCommand(log logger.StructuredLogger, validate *validator.Validate, conf *config.Config) error {
	metricsRegisterer := prometheus.DefaultRegisterer
	metricsGatherer := prometheus.DefaultGatherer

	shutdownHandler := bootstrap.ShutdownHandler(
		log,
		conf.Shutdown,
	)

	metricsServer := bootstrap.MetricsServer(
		log,
		conf.Metrics,
		metricsRegisterer,
		metricsGatherer,
	)

	healthcheckServer := bootstrap.HealthcheckServer(
		log,
		conf.HealthCheck,
	)

	mainDB := bootstrap.MainDB(log, conf.DB, metricsRegisterer)
	if err := mainDB.Ping(); err != nil {
		log.Error(
			logMessageMainDBPingError,
			emojiField("❌"),
			logger.ErrorField(err),
		)

		return errors.Wrap(err, "main db ping")
	}
	defer mainDB.Close()

	if err := mainDB.Migrate(); err != nil {
		log.Error(
			logMessageMainDBMigrateError,
			emojiField("❌"),
			logger.ErrorField(err),
		)

		return errors.Wrap(err, "main db migrate")
	}

	ethereumClient := bootstrap.EthereumClient(log, conf.Ethereum)
	defer ethereumClient.Close()

	if err := ethereumClient.Connect(context.TODO()); err != nil {
		log.Error(
			logMessageEthereumClientConnectError,
			emojiField("❌"),
			logger.ErrorField(err),
		)

		return errors.Wrap(err, "ethereum client connect")
	}

	fetcher := bootstrap.Fetcher(
		log,
		mainDB,
		ethereumClient,
	)

	limeServer := bootstrap.LimeServer(
		log,
		validate,
		mainDB,
		fetcher,
		conf.Lime,
	)

	// Task Group

	taskGroup := task.NewGroup()

	log.Info(
		logMessageTaskGroupStart,
		emojiField("🚀"),
	)

	taskGroup.Go(
		shutdownHandler.Run,
		metricsServer.Run,
		healthcheckServer.Run,
		limeServer.Run,
	)

	err := taskGroup.Wait(context.TODO())

	if err != nil && !errors.Is(err, shutdownhandler.ErrShutdown) {
		log.Error(
			logMessageTaskGroupError,
			emojiField("❌"),
			logger.ErrorField(err),
		)

		return errors.Wrap(err, "task group exits with error")
	}

	log.Info(
		logMessageTaskGroupShutdown,
		emojiField("🛑"),
	)

	return nil
}
