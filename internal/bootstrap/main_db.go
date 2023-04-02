package bootstrap

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	"github.com/ju-popov/the-ethereum-fetcher/internal/db/maindb"
	"github.com/ju-popov/the-ethereum-fetcher/internal/db/postgresql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sumup-oss/go-pkgs/logger"
)

func MainDB(
	log logger.StructuredLogger,
	conf *config.DB,
	metricsRegisterer prometheus.Registerer,
) *maindb.Client {
	return maindb.NewClient(
		log,
		*conf.Name,
		postgresql.NewBuilder().
			WithPort(*conf.PostgreSQL.Port).
			WithSSLMode(*conf.PostgreSQL.SSLMode).
			WithSchema(*conf.PostgreSQL.Schema).
			WithTimezone(*conf.PostgreSQL.Timezone).
			WithConnectTimeoutSeconds(*conf.PostgreSQL.ConnectTimeoutSeconds).
			WithMaxIdleConnections(*conf.PostgreSQL.MaxIdleConnections).
			WithMaxOpenConnections(*conf.PostgreSQL.MaxOpenConnections).
			WithConnectionMaxLifetime(*conf.PostgreSQL.ConnectionMaxLifetime).
			WithConnectionMaxIdleTime(*conf.PostgreSQL.ConnectionMaxIdleTime).
			WithMetricsRegisterer(metricsRegisterer).
			Build(
				log,
				*conf.Name,
				*conf.PostgreSQL.Host,
				*conf.PostgreSQL.Username,
				*conf.PostgreSQL.Password,
				*conf.PostgreSQL.Database,
			),
	)
}
