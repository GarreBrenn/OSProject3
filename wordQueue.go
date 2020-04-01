package main

import "fmt"
import "os"
import "strconv"

var numTasks int

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("USAGE: Wrong number of arguments.")
		fmt.Println("Please enter the (int) number of consumer tasks to run.\n")
	}

	if num, err := strconv.Atoi(args[0]); err == nil {
		numTasks = num
		fmt.Println(numTasks + 5)
	} else {
		fmt.Println("USAGE: Wrong argument type.")
		fmt.Println("Please enter the (int) number of consumer tasks to run.\n")
	}
}
