package main

import (
	"fmt"
)

func main() {
	f := func(a []string) ([]string, []string) {
		return a[2:], a[0:2]
	}
	m := []string{
		"one", "two", "three", "four", "five", "six",
	}
	s := []string{}
	fmt.Println(m)
	for len(m) > 0 {
		m, s = f(m)
		fmt.Println(s, " -> ", m)
	}
}
