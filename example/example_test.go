// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package example

import (
	"strings"
	"testing"

	"github.com/devnev/fakery"
	"github.com/devnev/fakery/internal/testutils"
	"github.com/google/go-cmp/cmp"
)

//go:generate go run github.com/devnev/fakery/cmd/fakery -src example_types.go -name ToBeMocked -dst example_mock.go

func TestFakeryMatch(t *testing.T) {
	m := &Mock_ToBeMocked{}
	On_ToBeMocked_Get(m, fakery.Equal("hello"), fakery.Returning1(Returned(returned("hello"))))
	var i ToBeMocked = m
	i.Get("hello").Hello()
}

func TestFakeryMatchNilInterface(t *testing.T) {
	m := &Mock_ToBeMocked{}
	On_ToBeMocked_Init(m, fakery.Equal[Required](nil), fakery.Any[string](), fakery.ReturningNothing())
	var i ToBeMocked = m
	i.Init(nil, "hello")
}

func TestFakeryNoMatch(t *testing.T) {
	m := &Mock_ToBeMocked{}
	On_ToBeMocked_Get(m, fakery.Equal("hello"), fakery.Returning1(Returned(returned("hello"))))
	On_ToBeMocked_Get(m, fakery.Equal("goodbye"), fakery.Returning1(Returned(returned("hello"))), fakery.Times(0))
	var i ToBeMocked = m

	out := testutils.RecordStderr(t)
	pv := testutils.CapturePanic(func() {
		i.Get("goodbye").Hello()
	})
	<-out.Stop()

	const expectedPanic = "no match for call to Get"
	if pv != expectedPanic {
		t.Errorf("expected pv %q, got %#v", expectedPanic, pv)
	}

	deterministicOutput := strings.ReplaceAll(out.Out(), "\u00a0", " ")
	expectdOutput := strings.Join([]string{
		`Matcher 1 (` + testutils.FileLine(-17) + `)`,
		"	Arg 0:",
		"		  string(",
		"		- \t\"hello\",",
		"		+ \t\"goodbye\",",
		"		  )",
		`Matcher 2 (` + testutils.FileLine(-22) + `)`,
		"	Times(0):",
		"		Called 0 times out of 0",
		"",
	}, "\n")
	if d := cmp.Diff(expectdOutput, deterministicOutput); d != "" {
		t.Errorf("output differed: %s", d)
	}
}

func TestFakeryMatchOnce(t *testing.T) {
	m := &Mock_ToBeMocked{}
	On_ToBeMocked_Get(m, fakery.Equal("hello"), fakery.Returning1(Returned(returned("hello"))), fakery.Once())
	var i ToBeMocked = m
	i.Get("hello").Hello()

	out := testutils.RecordStderr(t)
	pv := testutils.CapturePanic(func() {
		i.Get("hello").Hello()
	})
	<-out.Stop()

	const expectedPanic = "no match for call to Get"
	if pv != expectedPanic {
		t.Errorf("expected pv %q, got %#v", expectedPanic, pv)
	}
	deterministicOutput := strings.ReplaceAll(out.Out(), "\u00a0", " ")
	expectdOutput := strings.Join([]string{
		`Matcher 1 (` + testutils.FileLine(-16) + `)`,
		"	Once():",
		"		Already called",
		"",
	}, "\n")
	if d := cmp.Diff(expectdOutput, deterministicOutput); d != "" {
		t.Errorf("output differed: %s", d)
	}
}

func TestFakeryWithParametrisedNothingReturner(t *testing.T) {
	m := &Mock_ToBeMocked{}
	On_ToBeMocked_Init(m, fakery.Any[Required](), fakery.Any[string](), func(r Required, s string) (string, func()) {
		return "", func() {}
	})
	var i ToBeMocked = m
	i.Init(nil, "hello")
}

func TestFakeryWithParametrisedValueReturner(t *testing.T) {
	m := &Mock_ToBeMocked{}
	On_ToBeMocked_Get(m, fakery.Any[string](), func(string) (string, func() Returned) {
		return "", func() Returned {
			return nil
		}
	})
	var i ToBeMocked = m
	i.Get("hello")
}

func TestFakeryWithReturnSequence(t *testing.T) {
	m := &Mock_ToBeMocked{}
	On_ToBeMocked_Get(m, fakery.Any[string](), fakery.Returning1v[Returned](
		returned("hello"),
		returned("goodbye"),
	))
	var i ToBeMocked = m
	r := i.Get("something")
	if e := returned("hello"); r != e {
		t.Errorf("expected %q, got %q", e, r)
	}
	r = i.Get("something")
	if e := returned("goodbye"); r != e {
		t.Errorf("expected %q, got %q", e, r)
	}
	out := testutils.RecordStderr(t)
	pv := testutils.CapturePanic(func() {
		i.Get("something")
	})
	<-out.Stop()
	if e := "no match for call to Get"; pv != e {
		t.Errorf("expected panic %q, got %#v", e, pv)
	}
	if o, e := out.Out(), "Matcher 1 ("+testutils.FileLine(-21)+")\n\tReturn sequence exhausted\n"; o != e {
		t.Errorf("expected output %q, got %#v", e, o)
	}
}

func TestFakeryVariadicEquals(t *testing.T) {
	m := &Mock_ToBeMocked{}
	On_ToBeMocked_Print(m, fakery.VarArgs(fakery.Equal("hello"), fakery.Equal("goodbye")), fakery.ReturningNothing())
	var i ToBeMocked = m

	out := testutils.RecordStderr(t)
	pv := testutils.CapturePanic(func() {
		i.Print("oh no")
	})
	<-out.Stop()

	const expectedPanic = "no match for call to Print"
	if pv != expectedPanic {
		t.Errorf("expected pv %q, got %#v", expectedPanic, pv)
	}

	deterministicOutput := strings.ReplaceAll(out.Out(), "\u00a0", " ")
	expectdOutput := strings.Join([]string{
		`Matcher 1 (` + testutils.FileLine(-16) + `)`,
		"	Arg 0:",
		"		  string(",
		"		- \t\"hello\",",
		"		+ \t\"oh no\",",
		"		  )",
		"	VarArgs:",
		"		No argument for vararg matcher 1 (argument 1)",
		"",
	}, "\n")
	if d := cmp.Diff(expectdOutput, deterministicOutput); d != "" {
		t.Errorf("output differed: %s", d)
	}
}

type returned string

func (r returned) Hello() {
	println(string(r))
}
