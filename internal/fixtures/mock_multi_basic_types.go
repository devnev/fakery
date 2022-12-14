package fixtures

import fakery_gendeps "github.com/devnev/fakery/gendeps"
import fakery "github.com/devnev/fakery"

type Mock_MultiBasicTypes struct {
	matchers fakery_gendeps.MatchSet
}

func (m *Mock_MultiBasicTypes) InNoneOutTwoAnon() (string, int) {
	ret := fakery_gendeps.Called(&m.matchers, "InNoneOutTwoAnon", []any{})
	r0, _ := ret[0].(string)
	r1, _ := ret[1].(int)
	return r0, r1
}

func (m *Mock_MultiBasicTypes) InTwoAnonOutNone(
	a0 string,
	a1 int,
) {
	fakery_gendeps.Called(&m.matchers, "InTwoAnonOutNone", []any{&a0, &a1})
}

func (m *Mock_MultiBasicTypes) InTwoAnonOutTwoAnon(
	a0 float32,
	a1 int,
) (string, bool) {
	ret := fakery_gendeps.Called(&m.matchers, "InTwoAnonOutTwoAnon", []any{&a0, &a1})
	r0, _ := ret[0].(string)
	r1, _ := ret[1].(bool)
	return r0, r1
}

func (m *Mock_MultiBasicTypes) InTwoCombinedOutNone(
	a0 float32,
	a1 float32,
) {
	fakery_gendeps.Called(&m.matchers, "InTwoCombinedOutNone", []any{&a0, &a1})
}

func (m *Mock_MultiBasicTypes) InTwoNamedOutNone(
	a0 int,
	a1 string,
) {
	fakery_gendeps.Called(&m.matchers, "InTwoNamedOutNone", []any{&a0, &a1})
}

func On_MultiBasicTypes_InNoneOutTwoAnon(
	m *Mock_MultiBasicTypes,
	r func() (string, func() (string, int)),
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNoneOutTwoAnon", []any{}, r, o)
}

func On_MultiBasicTypes_InTwoAnonOutNone[
	R interface {
		func() (string, func()) | func(string, int) (string, func())
	},
](
	m *Mock_MultiBasicTypes,
	a0 func(int, string) string,
	a1 func(int, int) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InTwoAnonOutNone", []any{a0, a1}, r, o)
}

func On_MultiBasicTypes_InTwoAnonOutTwoAnon[
	R interface {
		func() (string, func() (string, bool)) | func(float32, int) (string, func() (string, bool))
	},
](
	m *Mock_MultiBasicTypes,
	a0 func(int, float32) string,
	a1 func(int, int) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InTwoAnonOutTwoAnon", []any{a0, a1}, r, o)
}

func On_MultiBasicTypes_InTwoCombinedOutNone[
	R interface {
		func() (string, func()) | func(float32, float32) (string, func())
	},
](
	m *Mock_MultiBasicTypes,
	a0 func(int, float32) string,
	a1 func(int, float32) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InTwoCombinedOutNone", []any{a0, a1}, r, o)
}

func On_MultiBasicTypes_InTwoNamedOutNone[
	R interface {
		func() (string, func()) | func(int, string) (string, func())
	},
](
	m *Mock_MultiBasicTypes,
	a0 func(int, int) string,
	a1 func(int, string) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InTwoNamedOutNone", []any{a0, a1}, r, o)
}

var _ MultiBasicTypes = &Mock_MultiBasicTypes{}
