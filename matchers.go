// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package fakery

import (
	"github.com/google/go-cmp/cmp"
)

// Equal does a deep-equal comparison of the argument to the provided value.
func Equal[T any](v T) func(T) string {
	return func(a T) string {
		return cmp.Diff(v, a)
	}
}

// Any matches any value
func Any[T any]() func(T) string {
	return func(t T) string { return "" }
}
