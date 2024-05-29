package fastcaller

import (
	_ "unsafe"
)

//go:noescape
//go:linkname Caller1 runtime.callers
func Caller1(skip int, pc *uintptr, len, cap int) int
