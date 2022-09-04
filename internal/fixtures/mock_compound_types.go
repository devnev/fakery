package fixtures

import fakery_gendeps "github.com/devnev/fakery/gendeps"
import fakery "github.com/devnev/fakery"

type Mock_CompoundTypes struct {
	matchers fakery_gendeps.MatchSet
}

func (m *Mock_CompoundTypes) InMapOutNone(
	a0 map[string]int,
) {
	fakery_gendeps.Called(&m.matchers, "InMapOutNone", []any{&a0})
}

func (m *Mock_CompoundTypes) InNoneOutMap() map[string]int {
	ret := fakery_gendeps.Called(&m.matchers, "InNoneOutMap", []any{})
	r0, _ := ret[0].(map[string]int)
	return r0
}

func (m *Mock_CompoundTypes) InNoneOutPointer() *string {
	ret := fakery_gendeps.Called(&m.matchers, "InNoneOutPointer", []any{})
	r0, _ := ret[0].(*string)
	return r0
}

func (m *Mock_CompoundTypes) InNoneOutSlice() []float32 {
	ret := fakery_gendeps.Called(&m.matchers, "InNoneOutSlice", []any{})
	r0, _ := ret[0].([]float32)
	return r0
}

func (m *Mock_CompoundTypes) InPointerOutNone(
	a0 *int,
) {
	fakery_gendeps.Called(&m.matchers, "InPointerOutNone", []any{&a0})
}

func (m *Mock_CompoundTypes) InSliceOutNone(
	a0 []string,
) {
	fakery_gendeps.Called(&m.matchers, "InSliceOutNone", []any{&a0})
}

func On_CompoundTypes_InMapOutNone[
	R interface {
		func() (string, func()) | func(map[string]int) (string, func())
	},
](
	m *Mock_CompoundTypes,
	a0 func(map[string]int) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InMapOutNone", []any{a0}, r, o)
}

func On_CompoundTypes_InNoneOutMap(
	m *Mock_CompoundTypes,
	r func() (string, func() map[string]int),
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNoneOutMap", []any{}, r, o)
}

func On_CompoundTypes_InNoneOutPointer(
	m *Mock_CompoundTypes,
	r func() (string, func() *string),
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNoneOutPointer", []any{}, r, o)
}

func On_CompoundTypes_InNoneOutSlice(
	m *Mock_CompoundTypes,
	r func() (string, func() []float32),
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InNoneOutSlice", []any{}, r, o)
}

func On_CompoundTypes_InPointerOutNone[
	R interface {
		func() (string, func()) | func(*int) (string, func())
	},
](
	m *Mock_CompoundTypes,
	a0 func(*int) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InPointerOutNone", []any{a0}, r, o)
}

func On_CompoundTypes_InSliceOutNone[
	R interface {
		func() (string, func()) | func([]string) (string, func())
	},
](
	m *Mock_CompoundTypes,
	a0 func([]string) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InSliceOutNone", []any{a0}, r, o)
}

var _ CompoundTypes = &Mock_CompoundTypes{}
