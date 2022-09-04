package fixtures

import fakery_gendeps "github.com/devnev/fakery/gendeps"
import fakery "github.com/devnev/fakery"
import sort "sort"

type Mock_ImportedTypes struct {
	matchers fakery_gendeps.MatchSet
}

func (m *Mock_ImportedTypes) InIface(
	a0 sort.Interface,
) {
	fakery_gendeps.Called(&m.matchers, "InIface", []any{&a0})
}

func (m *Mock_ImportedTypes) OutIface() sort.Interface {
	ret := fakery_gendeps.Called(&m.matchers, "OutIface", []any{})
	r0, _ := ret[0].(sort.Interface)
	return r0
}

func On_ImportedTypes_InIface[
	R interface {
		func() (string, func()) | func(sort.Interface) (string, func())
	},
](
	m *Mock_ImportedTypes,
	a0 func(sort.Interface) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InIface", []any{a0}, r, o)
}

func On_ImportedTypes_OutIface(
	m *Mock_ImportedTypes,
	r func() (string, func() sort.Interface),
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "OutIface", []any{}, r, o)
}

var _ ImportedTypes = &Mock_ImportedTypes{}
