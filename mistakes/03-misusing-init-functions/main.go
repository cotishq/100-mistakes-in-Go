package main

import "fmt"


var a = func() int {
	fmt.Println("var")
	return 0
} ()

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("main")
}