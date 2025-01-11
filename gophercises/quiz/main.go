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

	csvFileName := flag.String("csv", "problems.csv", "A csv file for the quiz: question,answer")
	timeLimit := flag.Int("limit", 5, "time limit")
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

	correct := 0
	questions := parseCsv(data)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemloop:
	for i, p := range questions {
		fmt.Printf("Question #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d", correct, len(questions))
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
