package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1(args []string) {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Printf("\n%.8fs elapsed\n%s\n", time.Since(start).Seconds(), s)
}

func echo2(args []string) {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("\n%.8fs elapsed\n%s\n", time.Since(start).Seconds(), s)
}

func echo3(args []string) {
	start := time.Now()
	fmt.Printf("\n%.8fs elapsed\n%s\n", time.Since(start).Seconds(), strings.Join(args[1:], " "))
}

func main() {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)
}