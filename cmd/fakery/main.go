// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

type Opts struct {
	Src  string
	Name string
	Dst  string
}

func main() {
	var opts Opts
	flag.StringVar(&opts.Src, "src", "", "")
	flag.StringVar(&opts.Name, "name", "", "")
	flag.StringVar(&opts.Dst, "dst", "", "")
	flag.Parse()
	os.Exit(run(opts))
}

func run(opts Opts) int {
	ins := Parse(opts.Src, opts.Name)

	if strings.HasSuffix(opts.Dst, ".go") {
		if len(ins) > 1 {
			println("dst file with multiple interfaces")
			return 1
		}
		_, src := Generator{buf: &strings.Builder{}}.Gen(ins[0])
		if src != "" {
			os.WriteFile(opts.Dst, []byte(src), 0o644)
		}
		return 0
	}

	for _, in := range ins {
		dst, src := Generator{buf: &strings.Builder{}}.Gen(in)
		if src != "" {
			os.WriteFile(filepath.Join(opts.Dst, dst), []byte(src), 0o644)
		}
	}
	return 0
}
