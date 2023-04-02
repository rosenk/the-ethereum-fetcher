package server

import "time"

const (
	defaultListenAddress           = "0.0.0.0:2222"
	defaultReadHeaderTimeout       = 10 * time.Second
	defaultReadTimeout             = 10 * time.Second
	defaultWriteTimeout            = 10 * time.Second
	defaultGracefulShutdownTimeout = 15 * time.Second
)
