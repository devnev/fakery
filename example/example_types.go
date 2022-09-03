// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
