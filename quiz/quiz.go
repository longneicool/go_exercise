package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type qAnda struct {
	question string
	answer   string
}

func getQAndA(records [][]string) ([]qAnda, int) {
	qas := make([]qAnda, 0)
	numInvaludQuestions := 0
	for _, r := range records {
		if len(r) != 2 {
			numInvaludQuestions++
			continue
		}

		qas = append(qas, qAnda{r[0], r[1]})
	}

	return qas, numInvaludQuestions
}

func isFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func main() {
	filename := flag.String("f", "problems.csv", "the problem file")
	flag.Parse()

	filePath := *filename

	if !isFileExist(filePath) {
		log.Fatalf("The file %s does not exist!", filePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file %s !", filePath)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read from the file %s! ", err.Error())
	}

	qAnda, incorrectQuestion := getQAndA(records)
	inReader := bufio.NewReader(os.Stdin)
	for _, s := range qAnda {
		fmt.Printf("->%s:", s.question)
		inputAnswer, err := inReader.ReadString('\n')
		inputAnswer = strings.Replace(inputAnswer, "\n", "", -1)
		if err != nil {
			continue
		}

		if s.answer != inputAnswer {
			incorrectQuestion++
		}
	}

	fmt.Printf("Total problem:%d, incorrect problem: %d\n", len(records), incorrectQuestion)
}
