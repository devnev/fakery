// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import "testing"

func TestRun(t *testing.T) {
	exit := run(Opts{
		Src: "../../internal/fixtures/interface.go",
		Dst: "../../internal/fixtures",
	})
	if exit != 0 {
		t.Fatalf("non-zero exit %d", exit)
	}
}

func TestRunOnExample(t *testing.T) {
	exit := run(Opts{
		Src:  "../../example/example_types.go",
		Name: "ToBeMocked",
		Dst:  "../../example/example_mock.go",
	})
	if exit != 0 {
		t.Fatalf("non-zero exit %d", exit)
	}
}
