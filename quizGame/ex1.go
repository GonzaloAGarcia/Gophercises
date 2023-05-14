package ex1

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type dataFrame struct {
	problem    string
	answer     string
	userAnswer string
}

func Ex1(path string, timeOut int) {

	var correctCount int = 0
	var questionCount int = 0

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fmt.Printf("Opened CSV from %s. \n", path)

	problems, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Press enter to start.")
	fmt.Scanln()
	go func() {
		time.Sleep(time.Second * time.Duration(timeOut))
		fmt.Println("Time's up!")
		fmt.Printf("You got %d out of %d questions right.\n", correctCount, questionCount)
		os.Exit(0)
	}()

	for _, row := range problems {

		var i string

		df := dataFrame{
			problem: row[0],
			answer:  row[1],
		}

		fmt.Println(df.problem)
		questionCount = questionCount + 1
		fmt.Scan(&i)

		df.userAnswer = i
		if df.userAnswer == df.answer {
			fmt.Print("Your Answer was correct \n")
			correctCount = correctCount + 1
		} else {
			fmt.Print("Your Answer was not correct \n")
		}

	}

}
