package fixtures

import fakery_gendeps "github.com/devnev/fakery/gendeps"
import fakery "github.com/devnev/fakery"
import sort "sort"

type Mock_InterfaceTypes struct {
	matchers fakery_gendeps.MatchSet
}

func (m *Mock_InterfaceTypes) InAny(
	a0 any,
) {
	fakery_gendeps.Called(&m.matchers, "InAny", []any{&a0})
}

func (m *Mock_InterfaceTypes) InEmptyInterface(
	a0 interface{},
) {
	fakery_gendeps.Called(&m.matchers, "InEmptyInterface", []any{&a0})
}

func (m *Mock_InterfaceTypes) InImported(
	a0 sort.Interface,
) {
	fakery_gendeps.Called(&m.matchers, "InImported", []any{&a0})
}

func (m *Mock_InterfaceTypes) OutImported() sort.Interface {
	ret := fakery_gendeps.Called(&m.matchers, "OutImported", []any{})
	r0, _ := ret[0].(sort.Interface)
	return r0
}

func On_InterfaceTypes_InAny[
	R interface {
		func() (string, func()) | func(any) (string, func())
	},
](
	m *Mock_InterfaceTypes,
	a0 func(int, any) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InAny", []any{a0}, r, o)
}

func On_InterfaceTypes_InEmptyInterface[
	R interface {
		func() (string, func()) | func(interface{}) (string, func())
	},
](
	m *Mock_InterfaceTypes,
	a0 func(int, interface{}) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InEmptyInterface", []any{a0}, r, o)
}

func On_InterfaceTypes_InImported[
	R interface {
		func() (string, func()) | func(sort.Interface) (string, func())
	},
](
	m *Mock_InterfaceTypes,
	a0 func(int, sort.Interface) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InImported", []any{a0}, r, o)
}

func On_InterfaceTypes_OutImported(
	m *Mock_InterfaceTypes,
	r func() (string, func() sort.Interface),
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "OutImported", []any{}, r, o)
}

var _ InterfaceTypes = &Mock_InterfaceTypes{}
