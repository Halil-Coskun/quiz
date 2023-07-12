package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	// parse the inputt csv string or default problems.csv
	csvFile := flag.String("csv", "problems.csv", "a csv file")
	flag.Parse()

	// open csv file
	file, err := os.Open(*csvFile)

	if err != nil {
		exit(fmt.Sprintf("Error opening the csv file: %s \n", *csvFile))
	}
	// read csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Error parsing from the csv")
	}

	timer := time.NewTimer(5 * time.Second)

	problems := parseLines(lines)
	score := 0

	for index, problem := range problems {
		fmt.Printf("%v) %v \n", index+1, problem.question)
		ansCh := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanln(&userAnswer)
			ansCh <- userAnswer
		}()
		select {
		case <-timer.C:
			fmt.Printf("Your total score is: %v \n", score)
			return
		case answer := <-ansCh:
			if answer == problem.answer {
				score++
			} else {
				fmt.Println("incorrect!")
			}
		}
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
