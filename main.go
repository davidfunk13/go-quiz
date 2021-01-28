package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

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

	fmt.Println(parsed)

	exitProgram("shutting down")
}

func exitProgram(message string) {
	fmt.Println(message)
	os.Exit(1)
}
