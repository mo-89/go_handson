package main

import (
	"fmt"
)

func main() {
	m := []string{
		"one", "two", "three", "four",
	}
	fmt.Println(m)
	m = insert(m, "*", 2)
	m = insert(m, "*", 1)
	fmt.Println(m)
}

func insert(a []string, v string, p int) (s []string) {
	s = append(a, "")
	fmt.Println("1st")
	fmt.Println(s)
	s = append(s[:p+1], s[p:len(s)-1]...)
	fmt.Println("2nd")
	fmt.Println(s)
	s[p] = v
	fmt.Println("3rd")
	fmt.Println(s)
	return
}
