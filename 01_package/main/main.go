package main

import (
	"fmt"
	"github.com/dashracer/udemy_golang_training/01_package/stringutil"
)

func main() {
	fmt.Println(stringutil.MyName)
	fmt.Println(stringutil.Reverse(stringutil.MyName))
}
