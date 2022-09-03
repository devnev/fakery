package testutils

import (
	"runtime"
	"strconv"
)

func FileLine(off int) string {
	_, file, line, _ := runtime.Caller(1)
	return file + ":" + strconv.Itoa(line+off)
}
