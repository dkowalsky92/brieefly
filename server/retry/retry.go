package retry

import (
	"fmt"
	"math"
	"time"

	"github.com/dkowalsky/brieefly/err"
	"github.com/dkowalsky/brieefly/log"
)

// Task -
type Task func() *err.Error

// Options -
type Options struct {
	Attempts uint64
	Interval time.Duration
	ErrorPfx string
	ErrorSfx string
}

// DefaultOptions -
func DefaultOptions() *Options {
	return NewOptions(nil, nil, nil, nil)
}

// NewOptions -
func NewOptions(attempts *uint64, interval *time.Duration, errorPfx *string, errorSfx *string) *Options {
	o := &Options{}
	if interval != nil {
		o.Attempts = *attempts
	} else {
		o.Attempts = math.MaxUint64
	}
	if interval != nil {
		o.Interval = *interval
	} else {
		o.Interval = 3
	}
	if errorPfx != nil {
		o.ErrorPfx = *errorPfx
	} else {
		o.ErrorPfx = "Task error:"
	}
	if errorSfx != nil {
		o.ErrorSfx = *errorSfx
	} else {
		o.ErrorSfx = fmt.Sprintf("attempting retry in %d seconds.", o.Interval)
	}
	return o
}

func perform(task Task, o Options) *err.Error {
	err := task()
	if err != nil {
		log.Error(fmt.Sprintf("%s %s - %s", o.ErrorPfx, err.Err, o.ErrorSfx))
		return err
	}

	return nil
}

func options(o *Options) Options {
	if o != nil {
		return *o
	}

	return *NewOptions(nil, nil, nil, nil)
}

// Perform -
func Perform(attempt uint64, opt *Options, task Task) {
	o := options(opt)
	if o.Attempts == attempt {
		log.Warning(fmt.Sprintf("Task couldn't succeed after %d attempts.", attempt))
		return
	}
	err := perform(task, o)
	if err != nil {
		time.AfterFunc(opt.Interval*time.Second, func() {
			Perform(attempt+1, opt, task)
		})
	}
}

// PerformInfinite -
func PerformInfinite(opt *Options, task Task) {
	o := options(opt)
	err := perform(task, o)
	if err != nil {
		time.AfterFunc(opt.Interval*time.Second, func() {
			PerformInfinite(opt, task)
		})
	}
}
