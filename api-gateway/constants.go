package main

import "time"

const (
	httpTimeoutDuration     = time.Second * 5
	httpIdleTimeoutDuration = time.Second * 30
	httpPort                = ":8000"
)
