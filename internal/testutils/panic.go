package testutils

func CapturePanic(fn func()) (pv any) {
	defer func() {
		pv = recover()
	}()
	fn()
	return pv
}
