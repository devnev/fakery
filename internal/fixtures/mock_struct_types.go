package fixtures

import fakery_gendeps "github.com/devnev/fakery/gendeps"
import fakery "github.com/devnev/fakery"
import http "net/http"

type Mock_StructTypes struct {
	matchers fakery_gendeps.MatchSet
}

func (m *Mock_StructTypes) InAnonEmptyStruct(
	a0 struct{},
) {
	fakery_gendeps.Called(&m.matchers, "InAnonEmptyStruct", []any{&a0})
}

func (m *Mock_StructTypes) InImportedStruct(
	a0 http.Server,
) {
	fakery_gendeps.Called(&m.matchers, "InImportedStruct", []any{&a0})
}

func On_StructTypes_InAnonEmptyStruct[
	R interface {
		func() (string, func()) | func(struct{}) (string, func())
	},
](
	m *Mock_StructTypes,
	a0 func(int, struct{}) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InAnonEmptyStruct", []any{a0}, r, o)
}

func On_StructTypes_InImportedStruct[
	R interface {
		func() (string, func()) | func(http.Server) (string, func())
	},
](
	m *Mock_StructTypes,
	a0 func(int, http.Server) string,
	r R,
	o ...fakery.Option,
) {
	fakery_gendeps.Add(&m.matchers, "InImportedStruct", []any{a0}, r, o)
}

var _ StructTypes = &Mock_StructTypes{}
