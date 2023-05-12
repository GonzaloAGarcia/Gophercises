package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type dataFrame struct {
	problem    string
	answer     string
	userAnswer string
}

func main() {

	var correctCount int = 0
	var questionCount int = 0
	var path = flag.String("path", "problems.csv", "Set path to CSV file with questions.")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fmt.Printf("Opened CSV from %s. \n", *path)

	problems, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, row := range problems {

		var i string

		df := dataFrame{
			problem: row[0],
			answer:  row[1],
		}

		fmt.Println(df.problem)
		fmt.Scan(&i)

		df.userAnswer = i
		if df.userAnswer == df.answer {
			fmt.Print("Your Answer was correct \n")
			correctCount = correctCount + 1
			questionCount = questionCount + 1
		} else {
			fmt.Print("Your Answer was not correct \n")
			questionCount = questionCount + 1
		}

	}

	fmt.Printf("You got %d out of %d questions right.\n", correctCount, questionCount)
}
