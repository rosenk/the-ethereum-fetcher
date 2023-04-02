package controller

import "time"

const (
	defaultListenAddress           = "0.0.0.0:5000"
	defaultReadHeaderTimeout       = 10 * time.Second
	defaultReadTimeout             = 10 * time.Second
	defaultWriteTimeout            = 10 * time.Second
	defaultGracefulShutdownTimeout = 15 * time.Second
)
