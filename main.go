package main

import "fmt"

//go:noinline
func fint(x int) int {
	return x
}

//go:noinline
func fstring(x string) string {
	return x
}

//go:noinline
func fstruct(x s) s {
	return x
}

type s struct {
	x byte
	y string
}

func main() {
	fmt.Println(fint(7))

	fmt.Println(fstruct(s{x: 'a', y: "a string"}))

	fmt.Println(fstring("helloworld"))
}
