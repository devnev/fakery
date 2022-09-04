// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package backend

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

type Matcher struct {
	args  []reflect.Value
	retFn reflect.Value
	opts  []Option
	file  string
	line  int
}

type MatchSet map[string][]Matcher

func Add(set *MatchSet, method string, args []any, ret any, opts []Option) {
	if *set == nil {
		*set = make(MatchSet)
	}
	if reflect.ValueOf(ret).IsNil() {
		panic("nil ret")
	}
	_, file, line, _ := runtime.Caller(3) // 0=backend, 1=gendeps, 2=mock, 3=user
	var argsRv []reflect.Value
	for _, a := range args {
		argsRv = append(argsRv, reflect.ValueOf(a))
	}
	(*set)[method] = append((*set)[method], Matcher{
		args:  argsRv,
		retFn: reflect.ValueOf(ret),
		opts:  opts,
		file:  file,
		line:  line,
	})
}

func Called(ms *MatchSet, m string, as []any) []any {
	var arv []reflect.Value
	for _, a := range as {
		arv = append(arv, reflect.ValueOf(a).Elem())
	}
	var dss [][]string
	for _, m := range (*ms)[m] {
		if ds, rs := check(m, as, arv); len(ds) == 0 {
			return rs
		} else {
			dss = append(dss, ds)
		}
	}
	for i, ds := range dss {
		fmt.Fprintf(os.Stderr, "Matcher %d (%s:%d)\n", i+1, (*ms)[m][i].file, (*ms)[m][i].line)
		for _, s := range ds {
			if s != "" {
				fmt.Fprint(os.Stderr, "\t"+strings.ReplaceAll(s, "\n", "\n\t\t")+"\n")
			}
		}
	}
	panic("no match for call to " + m)
}

func check(m Matcher, ain []any, arv []reflect.Value) (ds []string, rs []any) {
	for _, o := range m.opts {
		if d := o.called(ain); d != "" {
			ds = append(ds, d)
		}
	}
	if len(ds) > 0 {
		return ds, nil
	}
	var ne bool
	for i := 0; i < len(arv); i++ {
		d := m.args[i].Call([]reflect.Value{arv[i]})[0].Interface().(string)
		if len(d) > 0 {
			d = "Arg " + strconv.Itoa(i) + ":\n" + strings.TrimRight(d, "\n")
		}
		ds = append(ds, d)
		ne = ne || d != ""
	}
	if ne {
		return ds, nil
	}
	ds = nil
	for _, o := range m.opts {
		if d := o.matched(ain); d != "" {
			ds = append(ds, d)
		}
	}
	if len(ds) > 0 {
		return ds, nil
	}
	r := m.retFn
	if r.Type().NumIn() == 0 {
		for _, rv := range r.Call(nil) {
			rs = append(rs, rv.Interface())
		}
	} else {
		for _, rv := range r.Call(arv) {
			rs = append(rs, rv.Interface())
		}
	}
	for _, o := range m.opts {
		o.returned(ain, rs)
	}
	return nil, rs
}
