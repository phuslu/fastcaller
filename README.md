# fastcaller - Fast runtime.Caller with unsafe

## Getting Started

Play on https://go.dev/play/p/64yMcCIHg7L

```go
package main

import "github.com/phuslu/fastcaller"

func caller(skip int) (file string, line int32, name string) {
	var pc uintptr
	fastcaller.Caller1(skip, &pc, 1, 1)
	return fastcaller.FileLineName(pc)
}

func main() {
	println(caller(1))
	println(caller(1))
}
```

## Performance
```
goos: linux
goarch: amd64
pkg: github.com/phuslu/fastcaller
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkFileLine
BenchmarkFileLine-4       	 7518823	       159.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkFileLineName
BenchmarkFileLineName-4   	 6687564	       178.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/phuslu/fastcaller	2.740s
```
