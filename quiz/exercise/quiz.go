package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {

	// Create a new CSV reader reading from stdin
	filename := flag.String("csv", "problems.csv", " file in the format of 'question,answer'")

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemsloop:
	// for each line
	for i, p := range problems {
		// print  the question
		fmt.Printf("Problem is #%d: %s = ", i, p.q)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			fmt.Printf("time out!\n")

			// fmt.Printf("Your scorced %d out of %d \n", correct, len(problems))
			// break for loop (not select)
			break problemsloop
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("Correct! !!")
				correct++
			}
		}

		// read the answer

		// if the answer is correct
		// count the correct answers

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
