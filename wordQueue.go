package main

import "fmt"
import "os"
import "strconv"

var numTasks int
const signal = "0123456789qwertyuioplkjhgfdsazxcvbnm"

func countWords(str string) {
	//accepts a line
	//count and return all the words in the line
}

func task(id int, jobs <-chan string, results chan<- int) {
	fmt.Println("I am id: ",id)
	for n := range jobs {
		fmt.Println(n);
		results <- 4
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("USAGE: Wrong number of arguments.")
		fmt.Println("Please enter the (int) number of consumer tasks to run.\n")
		os.Exit(0)
	}

	if num, err := strconv.Atoi(args[0]); err == nil {
		numTasks = num
		fmt.Println("numTasks: ", numTasks)
	} else {
		fmt.Println("USAGE: Wrong argument type.")
		fmt.Println("Please enter the (int) number of consumer tasks to run.\n")
		os.Exit(0)
	}

	//create buffered channels
	//use the number of lines to determine buffer size (or don't use a buffer, idk)
	jobs := make(chan string)
	results := make(chan int)

	//start tasks
	for i := 0; i < numTasks; i++ {
		go task(i, jobs, results)
		fmt.Println("Starting task: ", i)
	}

	//add everything to the queue
	//enqueue a special string to signal the end
}
