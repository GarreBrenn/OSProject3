package main

import "fmt"
import "os"
import "strconv"
import "strings"

func countWords(str string) int {
	//accepts a line
	//count and return all the words in the line
	s := strings.Split(str, " ")
	return len(s)
}

func task(id int, jobs <-chan string, results chan<- int) {

	for n := range jobs {
		results <- countWords(n)
	}

}

func main() {
	var numTasks int
	var numLines int

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

	example := [...]string{"This is the first line",
			"This is the second line of course",
			"Naturally this is the third line, what else",
			"It follows by extrapolation that this line represents the fourth in the continuum of lines"}
	numLines = len(example)

	//create buffered channels
	//use the number of lines to determine buffer size (or don't use a buffer, idk)
	jobs := make(chan string, numLines)
	results := make(chan int, numLines)

	//start tasks
	for i := 0; i < numTasks; i++ {
		go task(i, jobs, results)
	}

	//add everything to the queue
	for i := 0; i < numLines; i++ {
		jobs <- example[i]
	}
	close(jobs)

	for i := 0; i < numLines; i++ {
		fmt.Println(<-results)
	}

}
