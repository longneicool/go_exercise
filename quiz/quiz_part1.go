package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	csvFilename := flag.String("csv", "problem.csv", "a csv file that contains the questions and answers")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Println("Failed to open the file")
		return
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to readALl from csv file")
		return
	}

	qAndA := parseLines(lines)
	fmt.Println(qAndA)

	correct := 0
	index := 0
	for question, standard := range qAndA {
		fmt.Printf("%d # %s", index, question)
		var answer int
		fmt.Scanf("%d\n", &answer)

		index++

		if answer != standard {
			fmt.Printf("Answer is not correct. %d -- %d\n", standard, answer)
			continue
		}
		fmt.Println("Answer is correct\n")
		correct++
	}

	fmt.Println("number of correct answers:  ", correct)
}

func parseLines(s [][]string) map[string]int {
	mapQAndA := make(map[string]int)
	for _, line := range s {
		if len(line) != 2 {
			fmt.Println("The len of the line is not 2: ")
			continue
		}

		question := line[0]
		answer := line[1]
		a, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Println("Fail to transfer the string to int")
			continue
		}

		mapQAndA[question] = a
	}

	return mapQAndA
}
