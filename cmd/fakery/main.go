package main

import (
	"flag"
	"os"
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

	if opts.Dst != "" {
		if len(ins) > 1 {
			println("dst with multiple interfaces")
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
		println(dst, len(src))
		if src != "" {
			os.WriteFile(dst, []byte(src), 0o644)
		}
	}
	return 0
}
