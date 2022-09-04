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
	"sync"
)

type Matcher struct {
	args  []reflect.Value
	retFn reflect.Value
	opts  []Option
	file  string
	line  int
}

type MatchSet struct {
	mu      sync.Mutex
	methods map[string][]Matcher
}

func Add(set *MatchSet, method string, args []any, ret any, opts []Option) {
	set.mu.Lock()
	defer set.mu.Unlock()
	if set.methods == nil {
		set.methods = make(map[string][]Matcher)
	}
	if reflect.ValueOf(ret).IsNil() {
		panic("nil ret")
	}
	_, file, line, _ := runtime.Caller(3) // 0=backend, 1=gendeps, 2=mock, 3=user
	var argsRv []reflect.Value
	for _, a := range args {
		argsRv = append(argsRv, reflect.ValueOf(a))
	}
	set.methods[method] = append(set.methods[method], Matcher{
		args:  argsRv,
		retFn: reflect.ValueOf(ret),
		opts:  opts,
		file:  file,
		line:  line,
	})
}

func Called(ms *MatchSet, m string, as []any) []any {
	ms.mu.Lock()
	locked := true
	defer func() {
		if locked {
			ms.mu.Unlock()
		}
	}()

	var arv []reflect.Value
	for _, a := range as {
		arv = append(arv, reflect.ValueOf(a).Elem())
	}

	var dss [][]string
	for _, m := range ms.methods[m] {
		if ds, retFn, opts := check(m, as, arv); len(ds) == 0 {
			locked = false
			ms.mu.Unlock()
			for _, o := range opts {
				o.unlocked()
			}
			var rs []any
			for _, rv := range retFn.Call(nil) {
				rs = append(rs, rv.Interface())
			}
			return rs
		} else {
			dss = append(dss, ds)
		}
	}

	for i, ds := range dss {
		fmt.Fprintf(os.Stderr, "Matcher %d (%s:%d)\n", i+1, ms.methods[m][i].file, ms.methods[m][i].line)
		for _, s := range ds {
			if s != "" {
				fmt.Fprint(os.Stderr, "\t"+strings.ReplaceAll(s, "\n", "\n\t\t")+"\n")
			}
		}
	}
	panic("no match for call to " + m)
}

func check(m Matcher, ain []any, arv []reflect.Value) (ds []string, ret reflect.Value, opts []Option) {
	for _, o := range m.opts {
		if d := o.called(ain); d != "" {
			ds = append(ds, d)
		}
	}
	if len(ds) > 0 {
		return ds, reflect.Value{}, m.opts
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
		return ds, reflect.Value{}, m.opts
	}

	var rv []reflect.Value
	if m.retFn.Type().NumIn() == 0 {
		rv = m.retFn.Call(nil)
	} else {
		rv = m.retFn.Call(arv)
	}
	if d := rv[0].Interface().(string); len(d) > 0 {
		return []string{d}, reflect.Value{}, m.opts
	}

	ds = nil
	for _, o := range m.opts {
		if d := o.matched(ain); d != "" {
			ds = append(ds, d)
		}
	}
	if len(ds) > 0 {
		return ds, reflect.Value{}, m.opts
	}

	return nil, rv[1], m.opts
}
