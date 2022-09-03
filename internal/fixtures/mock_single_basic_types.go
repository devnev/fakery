package fixtures

import fakery_gendeps "github.com/devnev/fakery/gendeps"
import fakery "github.com/devnev/fakery"

type Mock_SingleBasicTypes struct {
	matchers fakery_gendeps.MatchSet
}

func (m *Mock_SingleBasicTypes) InAnonOutNone(
	a0 int,
) {
	fakery_gendeps.Called(m.matchers, "InAnonOutNone", []any{&a0})
}

func (m *Mock_SingleBasicTypes) InNamedOutAnon(
	a0 string,
) int {
	ret := fakery_gendeps.Called(m.matchers, "InNamedOutAnon", []any{&a0})
	r0, _ := ret[0].(int)
	return r0
}

func (m *Mock_SingleBasicTypes) InNamedOutNamed(
	a0 int,
) string {
	ret := fakery_gendeps.Called(m.matchers, "InNamedOutNamed", []any{&a0})
	r0, _ := ret[0].(string)
	return r0
}

func (m *Mock_SingleBasicTypes) InNamedOutNone(
	a0 int,
) {
	fakery_gendeps.Called(m.matchers, "InNamedOutNone", []any{&a0})
}

func (m *Mock_SingleBasicTypes) InNoneOutAnonBasic() int {
	ret := fakery_gendeps.Called(m.matchers, "InNoneOutAnonBasic", []any{})
	r0, _ := ret[0].(int)
	return r0
}

func (m *Mock_SingleBasicTypes) InNoneOutNone() {
	fakery_gendeps.Called(m.matchers, "InNoneOutNone", []any{})
}

func On_SingleBasicTypes_InAnonOutNone[
	R interface{ func() | func(int) },
](
	m *Mock_SingleBasicTypes,
	a0 func(int) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InAnonOutNone", []any{a0}, r, o)
}

func On_SingleBasicTypes_InNamedOutAnon[
	R interface{ func() int | func(string) int },
](
	m *Mock_SingleBasicTypes,
	a0 func(string) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNamedOutAnon", []any{a0}, r, o)
}

func On_SingleBasicTypes_InNamedOutNamed[
	R interface {
		func() string | func(int) string
	},
](
	m *Mock_SingleBasicTypes,
	a0 func(int) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNamedOutNamed", []any{a0}, r, o)
}

func On_SingleBasicTypes_InNamedOutNone[
	R interface{ func() | func(int) },
](
	m *Mock_SingleBasicTypes,
	a0 func(int) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNamedOutNone", []any{a0}, r, o)
}

func On_SingleBasicTypes_InNoneOutAnonBasic[
	R interface{ func() int },
](
	m *Mock_SingleBasicTypes,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNoneOutAnonBasic", []any{}, r, o)
}

func On_SingleBasicTypes_InNoneOutNone[
	R interface{ func() },
](
	m *Mock_SingleBasicTypes,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNoneOutNone", []any{}, r, o)
}

var _ SingleBasicTypes = &Mock_SingleBasicTypes{}
