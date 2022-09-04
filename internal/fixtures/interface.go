// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package fixtures

import "sort"

//go:generate go run github.com/devnev/fakery/cmd/fakery -src interface.go

//fakery:unstable
type SingleBasicTypes interface {
	InNoneOutNone()
	InNoneOutAnonBasic() int
	InNamedOutNone(num int)
	InAnonOutNone(int)
	InNamedOutAnon(name string) int
	InNamedOutNamed(num int) (name string)
}

//fakery:unstable
type MultiBasicTypes interface {
	InTwoAnonOutNone(string, int)
	InNoneOutTwoAnon() (string, int)
	InTwoAnonOutTwoAnon(float32, int) (string, bool)
	InTwoNamedOutNone(num int, name string)
	InTwoCombinedOutNone(x, y float32)
}

//fakery:unstable
type ImportedTypes interface {
	InIface(sort.Interface)
	OutIface() sort.Interface
}

//fakery:unstable
type VariadicMethods interface {
	InVariadicOutNone(...int)
	InOneAndVariadicOutNone(int, ...string)
}
