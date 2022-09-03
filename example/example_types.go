package example

type Input struct {
	Val int
}

type Returned interface {
	Hello()
}

type Required interface {
	Start()
}

type ToBeMocked interface {
	Init(req Required, prefix string)
	Get(name string) Returned
	Add(in Input)
}
