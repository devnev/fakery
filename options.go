package fakery

import (
	"github.com/devnev/fakery/internal/backend"
	"golang.org/x/exp/constraints"
)

type Option = backend.Option

func Once() Option {
	return backend.Once()
}

func Times(n uint) Option {
	return backend.Times(n)
}

func CaptureCount[T constraints.Integer](counter *T) Option {
	return backend.CaptureCount(counter)
}

func AppendArgs(to *[][]any) Option {
	return backend.AppendArgs(to)
}
