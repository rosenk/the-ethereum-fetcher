package postgresql

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	defaultPort                  = 5432
	defaultSSLMode               = "require"
	defaultSchema                = "public"
	defaultTimezone              = "UTC"
	defaultConnectTimeoutSeconds = 10
	defaultMaxIdleConnections    = 2
	defaultMaxOpenConnections    = 10
	defaultConnectionMaxLifetime = 24 * time.Hour
	defaultConnectionMaxIdleTime = 24 * time.Hour
)

var defaultMetricsRegisterer = prometheus.DefaultRegisterer //nolint:gochecknoglobals
