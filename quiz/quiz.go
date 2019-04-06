package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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

func isEqual(input string, expected string) bool {
	input = strings.ToLower(strings.Replace(input, " ", "", -1))
	expected = strings.ToLower(strings.Replace(expected, " ", "", -1))

	return strings.Compare(input, expected) == 0
}

func main() {
	filename := flag.String("f", "problems.csv", "the problem file")
	timeLimit := flag.Uint("limit", 30, "The time limit of quiz")
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

	for {
		fmt.Println("Ready to start the quiz ? y or no")
		input, _ := inReader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if input == "yes" {
			break
		}

		if input != "no" {
			fmt.Println("Please input yes or no")
		}
	}

	timer1 := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	answerCh := make(chan string)
	for _, s := range qAnda {
		fmt.Printf("->%s:", s.question)
		go func() {
			answerStr, _ := inReader.ReadString('\n')
			answerStr = strings.Replace(answerStr, "\n", "", -1)
			answerCh <- answerStr
		}()

		select {
		case answer := <-answerCh:
			if !isEqual(answer, s.answer) {
				incorrectQuestion++
			}
		case <-timer1.C:
			fmt.Println("Time out")
			fmt.Printf("Total problem:%d, incorrect problem: %d\n", len(records), incorrectQuestion)
			return
		}
	}

	fmt.Printf("Total problem:%d, incorrect problem: %d\n", len(records), incorrectQuestion)
}
