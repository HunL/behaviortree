package main

import (
	"bufio"
	. "fmt"
	"os"
)

type Item struct {
	isFor bool
	forArgs []int
	args map[string]int
}

type board struct {
	items map[int]Item
}

var (
	i int
	j string
)

func main() {
	Println("behavior tree main function.")
	input2()
}

func input1() {
	Println("please input something: ")

	Scanln(&i, &j)
	Println("what you input is:", i, ", and ", j)
}

func input2() {
	inputReader := bufio.NewReader(os.Stdin)
	Println("Please input your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		Println("error occurs, exiting...")
		return
	}

	Printf("Your name is %s", input)
	switch input {
	// case "jeff\r\n": fallthrough
	case "jeff\n": fallthrough
	case "jack\r\n": fallthrough
	case "rose\r\n": Printf("Welcome %s\n", input)
	default: Printf("You are not welcome here! Goodbye!\n")
	}
}
