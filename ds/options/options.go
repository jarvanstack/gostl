package options

import (
	"sync"

	"github.com/dengjiawen8955/gostl/util/gosync"
)

// Options holds the Set's options
type Options struct {
	Locker gosync.Locker
}

// Option is a function  type used to set Options
type Option func(option *Options)

func WithSync() Option {
	return func(option *Options) {
		option.Locker = &sync.RWMutex{}
	}
}
