package fastcaller

import (
	"runtime"
	"testing"
	_ "unsafe"
)

func TestFileLine(t *testing.T) {
	var pc uintptr
	Caller1(0, &pc, 1, 1)
	file1, line1 := FileLine(pc)
	_, file2, line2, _ := runtime.Caller(0)

	if file1 != file2 {
		t.Errorf("FileLine file error: %q != %q", file1, file2)
	}

	if int(line1)+1 != line2 && int(line1)+2 != line2 {
		t.Errorf("FileLine line error: %d+2 != %d", line1, line2)
	}
}

func TestFileLineName(t *testing.T) {
	var pc uintptr
	Caller1(0, &pc, 1, 1)
	file1, line1, name1 := FileLineName(pc)
	t.Log(name1, file1, line1)
	pc, file2, line2, _ := runtime.Caller(0)
	frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()
	name2 := frame.Function
	t.Log(name2, file2, line2)

	if name1 != name2 {
		t.Errorf("FileLine file error: %q != %q", name1, name2)
	}

	if file1 != file2 {
		t.Errorf("FileLine file error: %q != %q", file1, file2)
	}

	if int(line1)+2 != line2 && int(line1)+3 != line2 {
		t.Errorf("FileLine line error: %d+3 != %d", line1, line2)
	}
}

func BenchmarkRuntimeFileLine(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		runtime.Caller(1)
	}
}

func BenchmarkRuntimeFileLineName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pc, _, _, _ := runtime.Caller(1)
		runtime.FuncForPC(pc)
	}
}

func BenchmarkFastcallerFileLine(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var pc uintptr
		Caller1(1, &pc, 1, 1)
		FileLine(pc)
	}
}

func BenchmarkFastcallerFileLineName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var pc uintptr
		Caller1(1, &pc, 1, 1)
		FileLineName(pc)
	}
}
