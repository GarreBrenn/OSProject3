/* Project 3
Garrett Brenner and Noah Giltmier

Written in "go"

"go run wordQueue.go [number of threads] < [stdin.txt]"
*/
package main

import "fmt"
import "os"
import "strconv"
import "strings"
import "bufio"
import "log"


// splits strings by " " (spaces) and then returns the word count of the line
func countWords(str string) int {
	//accepts a line
	//count and return all the words in the line
	s := strings.Split(str, " ")
	return len(s)
}

// the task that each thread will do 
func task(id int, jobs <-chan string, results chan<- int) {

	for n := range jobs {
		results <- countWords(n)
	}

}

//Main function
func main() {
	var numTasks int // number of threads
	var numLines int // number of lines in the stdin input


	// user error catch 
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("USAGE: Wrong number of arguments.")
		fmt.Println("Please enter the (int) number of consumer tasks to run.\n")
		os.Exit(0)
	}

	// prints the number of threads
	if num, err := strconv.Atoi(args[0]); err == nil {
		numTasks = num
		fmt.Println("numTasks: ", numTasks)
	
	} else { // more user error catching
		fmt.Println("USAGE: Wrong argument type.")
		fmt.Println("Please enter the (int) number of consumer tasks to run.\n")
		os.Exit(0)
	}

	
	var list []string // slice of lines from the user input

	// reads from a text file and seperates by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		list = append(list, scanner.Text())
	}

	// input error catching
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	numLines = len(list) // length of list aka number of lines in the input
	

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
		jobs <- list[i]
	}
	close(jobs)

	// print results
	for i := 0; i < numLines; i++ {
		fmt.Println(<-results)
	}

}
