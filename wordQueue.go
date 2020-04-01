package main

import "fmt"
import "os"


func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("USAGE: Wrong number of arguments.")
		fmt.Println("Please enter the number of consumer tasks to run.\n")
	} else {
		fmt.Println("Congrats!")
	}
}
