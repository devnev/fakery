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
func TestThing(t *testing.T) {
    mockIface := &Mock_IFace{}
    On_Iface_DoThing(mockIface, fakery.Equal("hello"), fakery.ReturnNothing, fakery.Once())
    // A *MockIface can be used anywhere an Iface is expected
    var realIface Iface  = mockIface
    realIface.DoThing("hello")
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
