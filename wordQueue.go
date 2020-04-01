package main

import "fmt"
import "os"
import "strconv"
import "strings"

var numTasks int
const signal = "0123456789qwertyuioplkjhgfdsazxcvbnm"

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
	}

	//add everything to the queue
	//enqueue a special string to signal the end
	example := [...]string{"This is the first line",
			"This is the second line of course",
			"Naturally this is the third line, what else",
			"It follows by extrapolation that this line represents the fourth in the continuum of lines"}

	for i := 0; i < len(example); i++ {
		jobs <- example[i]
	}
	close(jobs)

	for i := range results {
		fmt.Println(i)
	}

}
