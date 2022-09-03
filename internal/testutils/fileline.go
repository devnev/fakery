// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package testutils

import (
	"runtime"
	"strconv"
)

func FileLine(off int) string {
	_, file, line, _ := runtime.Caller(1)
	return file + ":" + strconv.Itoa(line+off)
}
