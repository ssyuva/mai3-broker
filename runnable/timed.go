package runnable

import (
	"context"
	"sync"
	"time"
)

type Timed struct {
	HWM int

	trgr chan interface{}
	once sync.Once
}

func NewTimed(hwm int) *Timed {
	return &Timed{
		HWM:  hwm,
		trgr: make(chan interface{}, hwm),
	}
}

func (t *Timed) Trigger(p interface{}) error {
	if len(t.trgr) == t.HWM {
		return ErrSystemBusy
	}
	t.trgr <- p
	return nil
}

func (t *Timed) Run(ctx context.Context, intvl time.Duration, f func()) (err error) {
	err = ErrSystemRunning
	t.once.Do(func() {
		err = nil
		var (
			before  time.Time
			elasped time.Duration
			timer   = time.NewTimer(0)
		)
		for {
			select {
			case <-t.trgr:
			case <-timer.C:
			case <-ctx.Done():
				err = ErrSystemInterrupted
				return
			}
			before = time.Now()
			f()
			elasped = time.Since(before)
			if elasped >= intvl {
				timer.Reset(0)
			} else {
				timer.Reset(intvl - elasped)
			}
		}
	})
	return
}
