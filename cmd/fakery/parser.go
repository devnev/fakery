// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"go/ast"
	"go/types"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

func Parse(path string, name string) []iface {
	conf := &packages.Config{
		Mode: packages.NeedTypes,
	}
	if name == "" {
		conf.Mode |= packages.NeedSyntax
	}
	var pkgs []*packages.Package
	if base := filepath.Base(path); strings.Contains(base, ".") && base != "..." && base != "." {
		pkgs, _ = packages.Load(conf, "file="+path)
	} else {
		pkgs, _ = packages.Load(conf, "pattern="+path)
	}
	if name != "" {
		return []iface{process(pkgs[0].Types, name)}
	}
	var ifaces []iface
	for _, pkg := range pkgs {
		var names []string
		for _, f := range pkg.Syntax {
			for _, d := range f.Decls {
				if g, _ := d.(*ast.GenDecl); g != nil {
					declFakery := false
					if g.Doc != nil {
						for _, c := range g.Doc.List {
							if strings.HasPrefix(c.Text, "//fakery:") {
								if c.Text != "//fakery:unstable" {
									panic("invalid fakery directive " + c.Text)
								}
								declFakery = true
							}
						}
					}
					for _, s := range g.Specs {
						if t, _ := s.(*ast.TypeSpec); t != nil {
							typeFakery := false
							for _, cg := range []*ast.CommentGroup{t.Comment, t.Doc} {
								if cg != nil {
									for _, c := range cg.List {
										if strings.HasPrefix(c.Text, "//fakery:") {
											if c.Text != "//fakery:unstable" {
												panic("invalid fakery directive " + c.Text)
											}
											typeFakery = true
										}
									}
								}
							}
							if declFakery || typeFakery {
								names = append(names, t.Name.Name)
							}
						}
					}
				}
			}
		}
		for _, name := range names {
			ifaces = append(ifaces, process(pkg.Types, name))
		}
	}
	return ifaces
}

func process(pkg *types.Package, name string) iface {
	target := pkg.Scope().Lookup(name).Type().(*types.Named)
	in := target.Underlying().(*types.Interface)
	iface := iface{
		pkgpath: pkg.Path(),
		pkg:     pkg.Name(),
		name:    target.Obj().Name(),
	}
	importSet := map[imp]struct{}{}
	for i, mc := 0, in.NumMethods(); i < mc; i++ {
		m := in.Method(i)
		sig := m.Type().(*types.Signature)
		iface.methods = append(iface.methods, method{
			name:       m.Name(),
			paramTypes: convertTypes(pkg, sig.Params(), sig.Variadic(), importSet),
			retTypes:   convertTypes(pkg, sig.Results(), false, importSet),
		})
	}
	for imp := range importSet {
		iface.imports = append(iface.imports, imp)
	}
	return iface
}

func convertTypes(pkg *types.Package, ts *types.Tuple, variadic bool, imps map[imp]struct{}) []string {
	var s []string
	for i, rc := 0, ts.Len(); i < rc; i++ {
		result := ts.At(i)
		if variadic && i == rc-1 {
			s = append(s, "..."+convertType(pkg, result.Type().(*types.Slice).Elem(), imps))
		} else {
			s = append(s, convertType(pkg, result.Type(), imps))
		}
	}
	return s
}

func convertType(pkg *types.Package, t types.Type, imps map[imp]struct{}) string {
	return types.TypeString(t, func(p *types.Package) string {
		if p == pkg {
			return ""
		}
		imps[imp{
			name: p.Name(),
			path: p.Path(),
		}] = struct{}{}
		return p.Name()
	})
}

type iface struct {
	pkgpath string
	pkg     string
	name    string
	imports []imp
	methods []method
}

type imp struct {
	name string
	path string
}

type method struct {
	name       string
	paramTypes []string
	retTypes   []string
}
