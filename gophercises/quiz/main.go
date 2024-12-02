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

	csvFileName := flag.String("csv", "problems.csv", "A csv file for the quiz: question,answer")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit("Not possible to read the file")
	}

	csv := csv.NewReader(file)

	data, err := csv.ReadAll()
	if err != nil {
		exit("Not possible to read the data")
	}

	questions := parseCsv(data)
	score := checkAnswer(questions)

	fmt.Printf("You scored %d out of %d", score, len(questions))
}

func checkAnswer(questions []problem) int {
	correct := 0

	for i, p := range questions {
		fmt.Printf("Question #%d: %s = \n", i+1, p.q)

		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			correct++
		}
	}

	return correct
}

func parseCsv(data [][]string) []problem {
	problems := make([]problem, len(data))
	for i, v := range data {
		problems[i] = problem{
			q: v[0],
			a: strings.TrimSpace(v[1]),
		}
	}

	return problems
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
