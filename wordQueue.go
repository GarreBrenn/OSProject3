package main

import "fmt"
import "os"
import "strconv"

var numTasks int
const signal = "0123456789qwertyuioplkjhgfdsazxcvbnm"

func task(id int) {
	fmt.Println("id: ",id)
	//dequeue off the queue
	//check to see if the line is the special string
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("USAGE: Wrong number of arguments.")
		fmt.Println("Please enter the (int) number of consumer tasks to run.\n")
	}

	if num, err := strconv.Atoi(args[0]); err == nil {
		numTasks = num
		fmt.Println("numTasks: ", numTasks)
	} else {
		fmt.Println("USAGE: Wrong argument type.")
		fmt.Println("Please enter the (int) number of consumer tasks to run.\n")
	}

	//start tasks
	for i := 0; i < numTasks; i++ {
		go task(i)
		fmt.Println("asdf")
	}

	//add everything to the queue
	//enqueue a special string to signal the end
}
