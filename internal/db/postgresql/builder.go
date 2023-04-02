package postgresql

import (
	"database/sql"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/dlmiddlecote/sqlstats"
	"github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Builder struct {
	port                  int
	sslMode               string
	schema                string
	timezone              string
	connectTimeoutSeconds int
	maxIdleConnections    int
	maxOpenConnections    int
	connectionMaxLifetime time.Duration
	connectionMaxIdleTime time.Duration
	metricsRegisterer     prometheus.Registerer
}

func NewBuilder() *Builder {
	return &Builder{
		port:                  defaultPort,
		sslMode:               defaultSSLMode,
		schema:                defaultSchema,
		timezone:              defaultTimezone,
		connectTimeoutSeconds: defaultConnectTimeoutSeconds,
		maxIdleConnections:    defaultMaxIdleConnections,
		maxOpenConnections:    defaultMaxOpenConnections,
		connectionMaxLifetime: defaultConnectionMaxLifetime,
		connectionMaxIdleTime: defaultConnectionMaxIdleTime,
		metricsRegisterer:     defaultMetricsRegisterer,
	}
}

func (b *Builder) WithPort(port int) *Builder {
	b.port = port

	return b
}

func (b *Builder) WithSSLMode(sslMode string) *Builder {
	b.sslMode = sslMode

	return b
}

func (b *Builder) WithSchema(schema string) *Builder {
	b.schema = schema

	return b
}

func (b *Builder) WithTimezone(timezone string) *Builder {
	b.timezone = timezone

	return b
}

func (b *Builder) WithConnectTimeoutSeconds(connectTimeoutSeconds int) *Builder {
	b.connectTimeoutSeconds = connectTimeoutSeconds

	return b
}

func (b *Builder) WithMaxIdleConnections(maxIdleConnections int) *Builder {
	b.maxIdleConnections = maxIdleConnections

	return b
}

func (b *Builder) WithMaxOpenConnections(maxOpenConnections int) *Builder {
	b.maxOpenConnections = maxOpenConnections

	return b
}

func (b *Builder) WithConnectionMaxLifetime(connectionMaxLifetime time.Duration) *Builder {
	b.connectionMaxLifetime = connectionMaxLifetime

	return b
}

func (b *Builder) WithConnectionMaxIdleTime(connectionMaxIdleTime time.Duration) *Builder {
	b.connectionMaxIdleTime = connectionMaxIdleTime

	return b
}

func (b *Builder) WithMetricsRegisterer(metricsRegisterer prometheus.Registerer) *Builder {
	b.metricsRegisterer = metricsRegisterer

	return b
}

//nolint:funlen
func (b *Builder) Build(
	log logger.StructuredLogger,
	name string,
	host string,
	username string,
	password string,
	database string,
) *sql.DB {
	log = log.With(nameField(name))

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(username, password),
		Host:   net.JoinHostPort(host, strconv.Itoa(b.port)),
		Path:   database,
	}

	query := dsn.Query()
	query.Set("search_path", b.schema)
	query.Set("sslmode", b.sslMode)
	query.Set("connect_timeout", strconv.Itoa(b.connectTimeoutSeconds))
	query.Set("timezone", b.timezone)

	dsn.RawQuery = query.Encode()

	//nolint:varnamelen
	db := sqldblogger.OpenDriver(
		dsn.String(),
		&pq.Driver{},
		NewLoggerAdapter(log),
		sqldblogger.WithErrorFieldname("error_sql"),
		sqldblogger.WithDurationFieldname("duration"),
		sqldblogger.WithTimeFieldname("time_sql"),
		sqldblogger.WithSQLQueryFieldname("query"),
		sqldblogger.WithSQLArgsFieldname("args"),
		sqldblogger.WithMinimumLevel(sqldblogger.LevelDebug),
		sqldblogger.WithLogArguments(true),
		sqldblogger.WithDurationUnit(sqldblogger.DurationNanosecond),
		sqldblogger.WithTimeFormat(sqldblogger.TimeFormatRFC3339Nano),
		sqldblogger.WithLogDriverErrorSkip(false),
		sqldblogger.WithSQLQueryAsMessage(false),
		sqldblogger.WithConnectionIDFieldname("conn_id"),
		sqldblogger.WithStatementIDFieldname("stmt_id"),
		sqldblogger.WithTransactionIDFieldname("tx_id"),
		sqldblogger.WithWrapResult(true),
		sqldblogger.WithIncludeStartTime(true),
		sqldblogger.WithStartTimeFieldname("start"),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelInfo),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelInfo),
		sqldblogger.WithExecerLevel(sqldblogger.LevelInfo),
	)

	db.SetMaxIdleConns(b.maxIdleConnections)
	db.SetMaxOpenConns(b.maxOpenConnections)
	db.SetConnMaxLifetime(b.connectionMaxLifetime)
	db.SetConnMaxIdleTime(b.connectionMaxIdleTime)

	b.metricsRegisterer.MustRegister(
		sqlstats.NewStatsCollector(name, db),
	)

	return db
}
