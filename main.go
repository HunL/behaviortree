package main

import (
	"bufio"
	. "fmt"
	"os"
)

type SeqItem struct {
	i int
}

type SelItem struct {
	i int
}

type CirItem struct {
	n      int
	action actionStruct
}

type actionStruct struct {
	target int
}

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
	input()
}

func input() {
	inputReader := bufio.NewReader(os.Stdin)
	Println("Please input pic. '1' for sequence, '2' for select, '3' for circulation.")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		Println("error occurs, exiting...")
		return
	}

	Printf("Your input is %s", input)
	switch input {
	case "1\r\n":
		fallthrough
	case "1\n":
		Printf("Welcome %s, now to makeSeqItem.\n", input)
		makeSeqItem()
	case "2\r\n":
		fallthrough
	case "2\n":
		Printf("Welcome %s, now to makeSelItem.\n", input)
		makeSelItem()
	case "3\r\n":
		fallthrough
	case "3\n":
		Printf("Welcome %s, now to makeCirItem.\n", input)
		makeCirItem()
	default:
		Printf("Invalid input! Goodbye!\n")
	}
}

func makeSeqItem() {
	return
}
func makeSelItem() {
	return
}
func makeCirItem() {
	return
}

func input1() {
	Println("please input something: ")

	Scanln(&i, &j)
	Println("what you input is:", i, ", and ", j)
}
