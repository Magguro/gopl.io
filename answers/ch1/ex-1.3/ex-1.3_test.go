package main

import (
	"testing"
	"strings"
)

/*
❯ go test -bench=. -benchmem

goos: darwin
goarch: arm64
pkg: main/answers/ch1/ex-1.3
BenchmarkEcho1-8         7696626               143.2 ns/op            56 B/op          3 allocs/op
BenchmarkEcho2-8         7682931               144.0 ns/op            56 B/op          3 allocs/op
BenchmarkEcho3-8        18501446               64.22 ns/op           24 B/op          1 allocs/op
PASS
ok      main/answers/ch1/ex-1.3 4.113s
*/

var (
	args = []string{"exec", "arg0", "arg1", "arg2", "arg3"}
)

func test_echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	// fmt.Println(s)
}

func test_echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func test_echo3(args []string) {
	strings.Join(args[1:], " ")
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_echo3(args)
	}
}
