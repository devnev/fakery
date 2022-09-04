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
	On_ToBeMocked_Get(m, fakery.Equal("hello"), fakery.Returning1(Returned(returned{})))
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
	On_ToBeMocked_Get(m, fakery.Equal("hello"), fakery.Returning1(Returned(returned{})))
	On_ToBeMocked_Get(m, fakery.Equal("goodbye"), fakery.Returning1(Returned(returned{})), fakery.Times(0))
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
	On_ToBeMocked_Get(m, fakery.Equal("hello"), fakery.Returning1(Returned(returned{})), fakery.Once())
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

type returned struct{}

func (r returned) Hello() {
	println("hello")
}
