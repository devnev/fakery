package main

import "testing"

func TestRun(t *testing.T) {
	exit := run(Opts{
		Src: "../../internal/fixtures/interface.go",
		// Name: "SingleBasicTypes",
		// Dst:  "../../fixtures/mock_single_basic_types.go",
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
