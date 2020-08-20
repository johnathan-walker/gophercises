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
	q string
	a string
}

func main() {
	filename := flag.String("filename", "problems.csv", "the CSV file containing questions and answers")
	timeLimit := flag.Int("timeLimit", 3, "the default timelimit for the quiz in seconds")
	_ = timeLimit
	flag.Parse()

	file, fileErr := os.Open(*filename)
	if fileErr != nil {
		exit(fmt.Sprintf("Error opening file: %s", *filename), 1)
	}

	r := csv.NewReader(file)
	lines, linesErr := r.ReadAll()
	if linesErr != nil {
		exit("Failed to parse the provided CSV file", 1)
	}

	problems := parseLinesToProblems(lines)
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.q)

		select {
		case <-timer.C:
			finish(correct, len(lines))
		default:
			answerChan := make(chan string)
			go func() {
				var answer string
				fmt.Scanf("%s\n", &answer)
				answerChan <- answer
			}()

			select {
			case answer := <-answerChan:
				if answer == p.a {
					correct++
				}
			case <-timer.C:
				finish(correct, len(lines))
			}
		}
	}
	finish(correct, len(lines))
}

func finish(correct int, total int) {
	fmt.Printf("%d / %d\n", correct, total)
	os.Exit(0)
}
func exit(message string, code int) {
	fmt.Println(message)
	os.Exit(code)
}

func parseLinesToProblems(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}
