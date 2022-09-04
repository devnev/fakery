package fixtures

import fakery_gendeps "github.com/devnev/fakery/gendeps"
import fakery "github.com/devnev/fakery"

type Mock_VariadicMethods struct {
	matchers fakery_gendeps.MatchSet
}

func (m *Mock_VariadicMethods) InOneAndVariadicOutNone(
	a0 int,
	a1 ...string,
) {
	fakery_gendeps.Called(&m.matchers, "InOneAndVariadicOutNone", []any{&a0, &a1})
}

func (m *Mock_VariadicMethods) InVariadicOutNone(
	a0 ...int,
) {
	fakery_gendeps.Called(&m.matchers, "InVariadicOutNone", []any{&a0})
}

func On_VariadicMethods_InOneAndVariadicOutNone[
	R interface {
		func() (string, func()) | func(int, []string) (string, func())
	},
](
	m *Mock_VariadicMethods,
	a0 func(int) string,
	a1 func([]string) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InOneAndVariadicOutNone", []any{a0, a1}, r, o)
}

func On_VariadicMethods_InVariadicOutNone[
	R interface {
		func() (string, func()) | func([]int) (string, func())
	},
](
	m *Mock_VariadicMethods,
	a0 func([]int) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InVariadicOutNone", []any{a0}, r, o)
}

var _ VariadicMethods = &Mock_VariadicMethods{}
