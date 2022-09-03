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
type MultiBasictypes interface {
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
