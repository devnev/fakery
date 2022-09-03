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

func Called(ms MatchSet, m string, as []any) []any {
	return backend.Called(ms, m, as)
}
