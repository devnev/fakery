// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package backend

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Once() Option {
	seen := false
	return option{
		calledFn: func(a []any) string {
			if seen {
				return "Once():\n\tAlready called"
			}
			return ""
		},
		matchedFn: func([]any) string {
			seen = true
			return ""
		},
	}
}

func Times(n uint) Option {
	var seen uint
	return option{
		calledFn: func(a []any) string {
			if seen >= n {
				return fmt.Sprintf("Times(%d):\n\tCalled %d times out of %d", n, seen, n)
			}
			return ""
		},
		matchedFn: func([]any) string {
			seen++
			return ""
		},
	}
}

func Increment[T constraints.Integer](counter *T) Option {
	return option{matchedFn: func([]any) string {
		*counter++
		return ""
	}}
}

func AppendArgs(to *[][]any) Option {
	return option{matchedFn: func(a []any) string {
		(*to) = append((*to), a)
		return ""
	}}
}

func WaitFor[T any](signal <-chan T) Option {
	return option{unlockedFn: func() {
		<-signal
	}}
}

type Option interface {
	called([]any) string
	matched([]any) string
	unlocked()
}

type option struct {
	calledFn   func([]any) string
	matchedFn  func([]any) string
	unlockedFn func()
}

func (o option) called(a []any) string {
	if o.calledFn != nil {
		return o.calledFn(a)
	}
	return ""
}

func (o option) matched(a []any) string {
	if o.matchedFn != nil {
		return o.matchedFn(a)
	}
	return ""
}

func (o option) unlocked() {
	if o.unlockedFn != nil {
		o.unlockedFn()
	}
}
