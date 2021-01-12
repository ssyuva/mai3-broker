package runnable

import (
	"errors"
)

var (
	ErrSystemBusy        = errors.New("system busy")
	ErrSystemInterrupted = errors.New("system interrupted")
	ErrSystemRunning     = errors.New("system already running")
)
