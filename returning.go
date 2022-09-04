// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package fakery

// ReturningNothing creates a return handler that always matches and returns no values.
func ReturningNothing() func() (string, func()) {
	return func() (string, func()) {
		return "", func() {}
	}
}

// Returning1 creates a return handler that always matches and returns one value.
func Returning1[T any](a0 T) func() (string, func() T) {
	return func() (string, func() T) {
		return "", func() T {
			return a0
		}
	}
}

// Returning2 creates a return handler that always matches and returns two values.
func Returning2[T0 any, T1 any](a0 T0, a1 T1) func() (string, func() (T0, T1)) {
	return func() (string, func() (T0, T1)) {
		return "", func() (T0, T1) {
			return a0, a1
		}
	}
}

// Returning3 creates a return handler that always matches and returns three values.
func Returning3[T0 any, T1 any, T2 any](a0 T0, a1 T1, a2 T2) func() (string, func() (T0, T1, T2)) {
	return func() (string, func() (T0, T1, T2)) {
		return "", func() (T0, T1, T2) {
			return a0, a1, a2
		}
	}
}

// Returning4 creates a return handler that always matches and returns four values.
func Returning4[T0 any, T1 any, T2 any, T3 any](a0 T0, a1 T1, a2 T2, a3 T3) func() (string, func() (T0, T1, T2, T3)) {
	return func() (string, func() (T0, T1, T2, T3)) {
		return "", func() (T0, T1, T2, T3) {
			return a0, a1, a2, a3
		}
	}
}

// Returning5 creates a return handler that always matches and returns five values.
func Returning5[T0 any, T1 any, T2 any, T3 any, T4 any](a0 T0, a1 T1, a2 T2, a3 T3, a4 T4) func() (string, func() (T0, T1, T2, T3, T4)) {
	return func() (string, func() (T0, T1, T2, T3, T4)) {
		return "", func() (T0, T1, T2, T3, T4) {
			return a0, a1, a2, a3, a4
		}
	}
}

// Returning1v creates a return handler that returns one item in the
// sequence at a time and matches as long as their are items remaining.
func Returning1v[T0 any](a0 ...T0) func() (string, func() T0) {
	seen := 0
	return func() (string, func() T0) {
		curr := seen
		seen++
		if curr >= len(a0) {
			return "Return sequence exhausted", nil
		}
		return "", func() T0 {
			return a0[curr]
		}
	}
}
