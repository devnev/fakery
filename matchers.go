// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package fakery

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

// Equal does a deep-equal comparison of the argument to the provided value.
func Equal[T any](v T) func(int, T) string {
	return func(i int, a T) string {
		if diff := cmp.Diff(v, a); diff != "" {
			return strings.TrimSpace(fmt.Sprintf("Arg %d:\n\t%s", i, strings.ReplaceAll(diff, "\n", "\n\t")))
		}
		return ""
	}
}

// Any matches any value
func Any[T any]() func(int, T) string {
	return func(int, T) string { return "" }
}

// Match items of a slice, for use with variadic functions
func VarArgs[T any](matchers ...func(int, T) string) func(int, []T) string {
	return func(i int, a []T) string {
		var diffs []string
		var matched bool = true
		for j := 0; j < len(a) || j < len(matchers); j++ {
			if j >= len(a) {
				diffs = append(diffs, fmt.Sprintf("VarArgs:\n\tNo argument for vararg matcher %d (argument %d)", j, i+j))
			} else if j >= len(matchers) {
				diffs = append(diffs, fmt.Sprintf("VarArgs:\n\tNo vararg matcher for argument %d (vararg %d)", i+j, j))
			} else {
				diffs = append(diffs, strings.TrimSpace(matchers[j](i+j, a[j])))
			}
			matched = matched && diffs[len(diffs)-1] == ""
		}
		if matched {
			return ""
		}
		return strings.Join(diffs, "\n")
	}
}
