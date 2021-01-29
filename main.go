package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

//struct defining quiz problem
type problem struct {
	question string
	answer   string
}

// function takes in array of csv rows contained in an array, returns an arrray of problems from
// question struct
func printProblems(parsed [][]string) []problem {

	// use builtin to make a slice of problems whos length is that of the input,
	// which the function's
	var row = make([]problem, len(parsed))

	// print question struct
	for i, l := range parsed {
		// assign one question for each item to each position of the array
		row[i] = problem{
			// assign the item at the first index of the array representing a line as the question
			question: l[0],
			answer:   strings.TrimSpace(l[1]),
		}
	}

	return row
}

func main() {
	fmt.Println("Starting up!")

	// flag to provide --input command line argument to seperate file.
	input := flag.String("input", "quiz.csv", "The path to the file we're reading")

	// read input flag
	flag.Parse()

	// open file for reading if provided, quiz if default. function takes pointer to input variable.
	file, err := os.Open(*input)

	//handle any errors
	if err != nil {
		fmt.Println(err, "error")
	}

	// use csv package to read the file we provide.
	f := csv.NewReader(file)

	// returns array of lines from the csv file
	parsed, err := f.ReadAll()

	// handle any errors
	if err != nil {
		exitProgram("There was an error reading the quiz file!")
	}

	// parse array of lines from the csv file to build a question
	problems := printProblems(parsed)

	fmt.Println(problems)

	correctanswers := 0

	for i, problem := range problems {
		//initializ string
		var answerinput string

		//print question.
		fmt.Printf("Problem #%d %s = \n", i+1, problem.question)

		// promp and read read answer input, assigns to
		// above initialized string
		fmt.Scanf("%s\n", &answerinput)

		if answerinput == problem.answer {
			correctanswers++
		}

		fmt.Printf("Correct answers: %d. \n", correctanswers)
	}

	exitProgram("shutting down")
}

func exitProgram(message string) {
	fmt.Println(message)
	os.Exit(1)
}
