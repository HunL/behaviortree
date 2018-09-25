package main

import (
	. "fmt"
)

var (
	i int
	j string
)

func main() {
	Println("behavior tree main function.")
	input1()
}

func input1() {
	Println("please input something: ")

	Scanln(&i, &j)
	Println("what you input is:", i, ", and ", j)
}
