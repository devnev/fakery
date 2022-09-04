// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package fakery

import (
	"github.com/google/go-cmp/cmp"
)

func Equal[T any](v T) func(T) string {
	return func(a T) string {
		return cmp.Diff(v, a)
	}
}

func Any[T any]() func(T) string {
	return func(t T) string { return "" }
}

func ReturningNothing() func() (string, func()) {
	return func() (string, func()) {
		return "", func() {}
	}
}

func Returning1[T any](a0 T) func() (string, func() T) {
	return func() (string, func() T) {
		return "", func() T {
			return a0
		}
	}
}

func Returning2[T0 any, T1 any](a0 T0, a1 T1) func() (string, func() (T0, T1)) {
	return func() (string, func() (T0, T1)) {
		return "", func() (T0, T1) {
			return a0, a1
		}
	}
}

func Returning3[T0 any, T1 any, T2 any](a0 T0, a1 T1, a2 T2) func() (string, func() (T0, T1, T2)) {
	return func() (string, func() (T0, T1, T2)) {
		return "", func() (T0, T1, T2) {
			return a0, a1, a2
		}
	}
}

func Returning4[T0 any, T1 any, T2 any, T3 any](a0 T0, a1 T1, a2 T2, a3 T3) func() (string, func() (T0, T1, T2, T3)) {
	return func() (string, func() (T0, T1, T2, T3)) {
		return "", func() (T0, T1, T2, T3) {
			return a0, a1, a2, a3
		}
	}
}

func Returning5[T0 any, T1 any, T2 any, T3 any, T4 any](a0 T0, a1 T1, a2 T2, a3 T3, a4 T4) func() (string, func() (T0, T1, T2, T3, T4)) {
	return func() (string, func() (T0, T1, T2, T3, T4)) {
		return "", func() (T0, T1, T2, T3, T4) {
			return a0, a1, a2, a3, a4
		}
	}
}
