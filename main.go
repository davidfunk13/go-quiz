package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func printProblems(parsed [][]string) string {
	for i, l := range parsed {
		fmt.Println(i)
		fmt.Println(l)
	}
	return "Ass"
}

func main() {
	fmt.Println("Starting up!")

	input := flag.String("input", "quiz.csv", "The path to the file we're reading")

	flag.Parse()

	file, err := os.Open(*input)

	if err != nil {
		fmt.Println(err, "error")
	}

	f := csv.NewReader(file)

	parsed, err := f.ReadAll()

	if err != nil {
		exitProgram("There was an error reading the quiz file!")
	}

	printProblems(parsed)
	fmt.Println(parsed)

	exitProgram("shutting down")
}

func exitProgram(message string) {
	fmt.Println(message)
	os.Exit(1)
}
