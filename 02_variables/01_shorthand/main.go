package main

import "fmt"

func main() {
	/*
		Note
		- Shorthand variable declaration using := with value, only in function
		- %T => variable type
		- %v => variable value
	 */
	a := 21
	b := "Addin Learn Golang"
	c := 3.14
	d := true

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
}