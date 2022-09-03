// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
