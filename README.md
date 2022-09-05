# Fakery - Mock codegen for easy fakes in Go

**Fake it 'til you make it!**

- Focused on fakes, not assertions
- Making use of generics for type-safety
- Keeping the size of generated code small
- Producing clear output when a call is not matched

## Usage

### Generating mocks

The CLI is _very_ hacked-together and subject to change. For now:

```sh
$ go install github.com/devnev/fakery/cmd/fakery@latest
# a specific interface
$ fakery -src src_file.go -name Iface -dst mock_iface.go
# all interfaces annotated with //fakery:unstable
$ fakery -src .
```

### Using generated mocks

```go
func TestDoThing(t *testing.T) {
    mockIface := &Mock_IFace{}
    On_Iface_DoThing(mockIface, fakery.Equal("hello"), fakery.ReturnNothing, fakery.Once())
    // A *MockIface can be used anywhere an Iface is expected
    var realIface Iface = mockIface
    realIface.DoThing("hello")
}

func TestDoThingCallArgs(t *testing.T) {
    mockIface := &Mock_IFace{}
	var doThingArgs [][]any
    On_Iface_DoThing(mockIface, fakery.Any[string](), fakery.ReturnNothing, fakery.AppendArgs(&doThingArgs))
	runTest(mockIface)
	// Use your preferred way of asserting, e.g. hand-rolled or testify assertions
	assert.Equal(t, [][]any{{"hello"}, {"world"}}, doThingArgs)
}
```

### Example failure

```
Matcher 1 (/gopath/github.com/devnev/fakery/example/example_test.go:30)
	Arg 0:
		  string(
		- 	"hello",
		+ 	"goodbye",
		  )
Matcher 2 (/gopath/github.com/devnev/fakery/example/example_test.go:31)
	Times(2):
		Called 2 times out of 2
--- FAIL: TestFakeryNoMatch (0.00s)
panic: no match for call to Get [recovered]
	panic: no match for call to Get

goroutine 34 [running]:
testing.tRunner.func1.2({0x11d6b00, 0xc00010a790})
	/usr/local/Cellar/go/1.19/libexec/src/testing/testing.go:1396 +0x24e
testing.tRunner.func1()
	/usr/local/Cellar/go/1.19/libexec/src/testing/testing.go:1399 +0x39f
panic({0x11d6b00, 0xc00010a790})
	/usr/local/Cellar/go/1.19/libexec/src/runtime/panic.go:884 +0x212
github.com/devnev/fakery/internal/backend.Called(0xc000112050?, {0x1213231, 0x3}, {0xc00010a640, 0x1, 0x1})
	/gopath/github.com/devnev/fakery/internal/backend/mocking.go:64 +0x485
github.com/devnev/fakery/gendeps.Called(...)
	/gopath/github.com/devnev/fakery/gendeps/mocking.go:18
github.com/devnev/fakery/example.(*Mock_ToBeMocked).Get(0xc000106e88, {0x1213c75, 0x7})
	/gopath/github.com/devnev/fakery/example/example_mock.go:19 +0xcb
github.com/devnev/fakery/example.TestFakeryNoMatch(0xc000116820)
	/gopath/github.com/devnev/fakery/example/example_test.go:36 +0x1a5
testing.tRunner(0xc000116820, 0x1222a00)
	/usr/local/Cellar/go/1.19/libexec/src/testing/testing.go:1446 +0x10b
created by testing.(*T).Run
	/usr/local/Cellar/go/1.19/libexec/src/testing/testing.go:1493 +0x35f
FAIL	github.com/devnev/fakery/example	0.307s
FAIL
```

### Extended usage & adaptation

The package provides utilities for matching arguments,

```go
Equal(value)
Any()
```

returning values,

```go
ReturningNothing()
Returning1(value)
Returning2(value1, value2)
//etc.
Returning1v(value1, value2)
```

recording calls,

```go
Increment(&counter)
AppendArgs(&args)
```

and other controls,

```go
Once()
Times(n)
WaitFor(ch)
```

Any of the above can be accomplished and extended with the parameters to the
`On_*` functions:

#### Arguments matchers

Argument matchers have the signature

```go
func(int, ArgType) string
```

The first argument is the argument index - mainly for use in failure messages -
the second is argument value. A non-empty return value indicates that the
argument does not match, with the reason as the string value. The mock is locked
against further calls to any method while the argument matchers are run.

#### Return handlers

Return value handlers have two possible signatures and two stages:

```go
// with args
func(...ArgTypes...) (string, func() (...ReturnTypes...))
// or without args
func() (string, func() (...ReturnTypes...))
```

The first stage runs before the match is confirmed (while the mock is locked
against further calls to any method), and may return a non-empty string to
indicate that the match failed.

If the first stage returns an empty string, it must also return a function for
the second stage. This stage runs when the call matching is complete and the
lock is released. It must return the values for the matched call to return.
