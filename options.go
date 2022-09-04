// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package fakery

import (
	"github.com/devnev/fakery/internal/backend"
	"golang.org/x/exp/constraints"
)

type Option = backend.Option

// Once changes a matcher to only match one invocation.
func Once() Option {
	return backend.Once()
}

// Times changes a matcher to only match n invocations.
func Times(n uint) Option {
	return backend.Times(n)
}

// Increment increments the counter when an invocation matches.
func Increment[T constraints.Integer](counter *T) Option {
	return backend.Increment(counter)
}

// AppendArgs appends the arguments of a matched call to a slice of argument slices.
func AppendArgs(to *[][]any) Option {
	return backend.AppendArgs(to)
}

// WaitFor waits for a signal before returning from a matched call. The delay
// occurs before the second stage of the return handler is invoked.
func WaitFor[T any](signal <-chan T) Option {
	return backend.WaitFor(signal)
}
