package example

import fakery_gendeps "github.com/devnev/fakery/gendeps"
import fakery "github.com/devnev/fakery"

type Mock_ToBeMocked struct {
	matchers fakery_gendeps.MatchSet
}

func (m *Mock_ToBeMocked) Add(
	a0 Input,
) {
	fakery_gendeps.Called(&m.matchers, "Add", []any{&a0})
}

func (m *Mock_ToBeMocked) Get(
	a0 string,
) Returned {
	ret := fakery_gendeps.Called(&m.matchers, "Get", []any{&a0})
	r0, _ := ret[0].(Returned)
	return r0
}

func (m *Mock_ToBeMocked) Init(
	a0 Required,
	a1 string,
) {
	fakery_gendeps.Called(&m.matchers, "Init", []any{&a0, &a1})
}

func (m *Mock_ToBeMocked) Print(
	a0 ...string,
) {
	fakery_gendeps.Called(&m.matchers, "Print", []any{&a0})
}

func On_ToBeMocked_Add[
	R interface {
		func() (string, func()) | func(Input) (string, func())
	},
](
	m *Mock_ToBeMocked,
	a0 func(int, Input) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "Add", []any{a0}, r, o)
}

func On_ToBeMocked_Get[
	R interface {
		func() (string, func() Returned) | func(string) (string, func() Returned)
	},
](
	m *Mock_ToBeMocked,
	a0 func(int, string) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "Get", []any{a0}, r, o)
}

func On_ToBeMocked_Init[
	R interface {
		func() (string, func()) | func(Required, string) (string, func())
	},
](
	m *Mock_ToBeMocked,
	a0 func(int, Required) string,
	a1 func(int, string) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "Init", []any{a0, a1}, r, o)
}

func On_ToBeMocked_Print[
	R interface {
		func() (string, func()) | func([]string) (string, func())
	},
](
	m *Mock_ToBeMocked,
	a0 func(int, []string) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "Print", []any{a0}, r, o)
}

var _ ToBeMocked = &Mock_ToBeMocked{}
