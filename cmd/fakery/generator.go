// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"go/format"
	"regexp"
	"strconv"
	"strings"
)

type Generator struct {
	buf *strings.Builder
}

func cameltosnake(s string) string {
	s = regexp.MustCompile(`^[A-Z]`).ReplaceAllStringFunc(s, func(s string) string {
		return strings.ToLower(s)
	})
	s = regexp.MustCompile(`[A-Z][a-z]`).ReplaceAllStringFunc(s, func(s string) string {
		return "_" + s
	})
	s = regexp.MustCompile(`[a-z][A-Z]`).ReplaceAllStringFunc(s, func(s string) string {
		return s[0:1] + "_" + s[1:2]
	})
	s = strings.ToLower(s)
	return s
}

func (g Generator) Gen(in iface) (string, string) {
	g.print("package ", in.pkg, "\n")
	g.print("import fakery_gendeps ", strconv.Quote("github.com/devnev/fakery/gendeps"))
	g.print("import fakery ", strconv.Quote("github.com/devnev/fakery"))
	for _, imp := range in.imports {
		g.print("import ", imp.name, " ", strconv.Quote(imp.path))
	}
	g.print("type Mock_", in.name, " struct {")
	g.print("matchers fakery_gendeps.MatchSet")
	g.print("}")

	for _, m := range in.methods {
		g.print("func (m *Mock_", in.name, ") ", m.name, "(")
		for i, pt := range m.paramTypes {
			g.print("a", strconv.Itoa(i), " ", pt, ",")
		}
		if len(m.retTypes) > 0 {
			g.print(") (", strings.Join(m.retTypes, ", "), ") {")
			g.print(
				"ret := fakery_gendeps.Called(&m.matchers, ",
				strconv.Quote(m.name),
				", []any{",
				strings.Join(numbered("&a", len(m.paramTypes)), ", "),
				"})",
			)
			for i, rt := range m.retTypes {
				g.print("r", strconv.Itoa(i), ", _ := ret[", strconv.Itoa(i), "].(", rt, ")")
			}
			g.print("return ", strings.Join(numbered("r", len(m.retTypes)), ", "))
		} else {
			g.print(") {")
			g.print(
				"fakery_gendeps.Called(&m.matchers, ",
				strconv.Quote(m.name),
				", []any{",
				strings.Join(numbered("&a", len(m.paramTypes)), ", "),
				"})",
			)
		}
		g.print("}\n")
	}

	for _, m := range in.methods {
		for i := range m.paramTypes {
			if strings.HasPrefix(m.paramTypes[i], "...") {
				m.paramTypes[i] = "[]" + strings.TrimPrefix(m.paramTypes[i], "...")
			}
		}
		if len(m.paramTypes) > 0 {
			g.print("func On_", in.name, "_", m.name, "[")
			g.print("R interface { func() (string, func() (", strings.Join(m.retTypes, ", "), ")) | func(", strings.Join(m.paramTypes, ", "), ") (string, func() (", strings.Join(m.retTypes, ", "), ")) },")
			g.print("](")
		} else {
			g.print("func On_", in.name, "_", m.name, "(")
		}
		g.print("m *Mock_", in.name, ",")
		for i, pt := range m.paramTypes {
			g.print("a", strconv.Itoa(i), " func(", pt, ") string,")
		}
		if len(m.paramTypes) > 0 {
			g.print("r R,")
		} else {
			g.print("r func() (string, func() (", strings.Join(m.retTypes, ", "), ")),")
		}
		g.print("o ...fakery.Option,")
		g.print(") {")
		g.print(
			"fakery_gendeps.Add(&m.matchers, ",
			strconv.Quote(m.name),
			", []any{",
			strings.Join(numbered("a", len(m.paramTypes)), ", "),
			"}, r, o)",
		)
		g.print("}\n")
	}

	g.print("var _ ", in.name, " = & Mock_", in.name, "{}")

	src, err := format.Source([]byte(g.buf.String()))
	if err != nil {
		println(err.Error())
		for i, s := range strings.Split(g.buf.String(), "\n") {
			print(fmt.Sprintf("%4d: %s\n", i+1, s))
		}
		return "", ""
	}
	outpath := "mock_" + cameltosnake(in.name) + ".go"
	return outpath, string(src)
}

func (g Generator) print(ss ...string) {
	for _, s := range ss {
		g.buf.WriteString(s)
	}
	g.buf.WriteRune('\n')
}

func numbered(pfx string, n int) (s []string) {
	for i := 0; i < n; i++ {
		s = append(s, pfx+strconv.Itoa(i))
	}
	return s
}
