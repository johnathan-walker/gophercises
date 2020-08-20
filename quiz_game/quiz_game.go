package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	filename := flag.String("filename", "problems.csv", "the CSV file containing questions and answers.")
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
	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.q)

		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("%d / %d\n", correct, len(lines))
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
