package fastcaller_test

import (
	"github.com/phuslu/fastcaller"
)

func ExampleFileLine() {
	var pc uintptr
	fastcaller.Caller1(0, &pc, 1, 1)
	println(fastcaller.FileLine(pc))
}

func ExampleFileLineName() {
	var pc uintptr
	fastcaller.Caller1(0, &pc, 1, 1)
	println(fastcaller.FileLineName(pc))
}
