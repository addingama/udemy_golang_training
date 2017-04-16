package main

import "fmt"

func main() {
	/*
		Note
		- Zero value variable declaration without assignment always have
		default value
		- %T => variable type
		- %v => variable value
	 */
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
}