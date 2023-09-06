package main

import (
	"fmt"
	"head-first/kata-also-here/kata6kyu"
)

func main() {
	a := make(chan string)
	b := make(chan string)

	go func() {
		a <- "first"
		a <- "second"
		a <- "third"
		close(a)
		close(b)
	}()

	c := kata6kyu.Merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
