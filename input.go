package main

// Importing packages to run program.

import (

	// bufio creates a buffer of input and output of the machine.
	"bufio"
	"fmt"
	"os"
)

func main() {

	// get some info from standard input and assign it to buf
	buf := bufio.NewReader(os.Stdin)

	// print out to user
	fmt.Println("What is your age?: ")

	// declaring variables for program. ReadBytes is to handle what's returned.
	age, err := buf.ReadBytes('\n')

	// conditonal on if error exists and if not, print age
	if err != nil {
		fmt.Println(age)
	} else {
		fmt.Println(err)
	}
}
