package main

import (
	"bufio"
	. "fmt"
	"os"
)

type Item struct {
	isFor   bool
	forArgs []int
	args    map[string]int
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
	Println("Please input pic. '1' for sequence, '2' for select, '3' for circulation.")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		Println("error occurs, exiting...")
		return
	}

	Printf("Your input is %s", input)
	switch input {
	// case "jeff\r\n": fallthrough
	case "1\r\n":
		Printf("Welcome %s\n", input)
	case "2\r\n":
		Printf("Welcome %s\n", input)
	case "3\r\n":
		Printf("Welcome %s\n", input)
	// case "jeff\n": fallthrough
	// case "jack\r\n": fallthrough
	// case "rose\r\n": Printf("Welcome %s\n", input)
	default:
		Printf("Invalid input! Goodbye!\n")
	}
}
