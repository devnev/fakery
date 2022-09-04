// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// This package is for reference from generated mocks only, DO NOT USE
package gendeps

import (
	"github.com/devnev/fakery"
	"github.com/devnev/fakery/internal/backend"
)

type (
	Matcher  = backend.Matcher
	MatchSet = backend.MatchSet
)

func Add(set *MatchSet, method string, args []any, ret any, opts []fakery.Option) {
	backend.Add(set, method, args, ret, opts)
}

func Called(ms *MatchSet, m string, as []any) []any {
	return backend.Called(ms, m, as)
}
