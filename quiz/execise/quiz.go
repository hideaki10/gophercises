package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {

	// Create a new CSV reader reading from stdin
	filename := flag.String("csv", " ", " file in the format of 'question,answer'")
	flag.Parse()

	// open the csv file
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Filed open the csv file %s\n", *filename)
		os.Exit(1)
	}

	// read the csv file
	r := csv.NewReader(file)
	// parse the csv file
	line, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Fail to parse the csv file")
	}

	problems := parseLines(line)

	correct := 0

	// for each line
	for i, p := range problems {
		// print  the question
		fmt.Printf("Problem is #%d: %s = ", i, p.q)

		// read the answer
		var answer string
		fmt.Scanf("%s\n", &answer)

		// if the answer is correct
		// count the correct answers
		if answer == p.a {
			fmt.Println("Correct! !!")
			correct++
		}

	}

	fmt.Printf("Your scorced %d out of %d \n", correct, len(problems))

}

func parseLines(line [][]string) []problem {
	ret := make([]problem, len(line))
	for i, l := range line {
		ret[i] = problem{
			q: l[0],
			a: l[1],
		}
	}
	return ret
}
