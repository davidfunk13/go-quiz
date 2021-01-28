package main

import (
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

	fmt.Println(file)
	exitProgram("shutting down")
}

func exitProgram(message string) {
	fmt.Println(message)
	os.Exit(1)
}
