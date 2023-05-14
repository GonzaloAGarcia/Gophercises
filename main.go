package main

import (
	"flag"

	ex1 "gophercises.gonzalogarcia.de/quizGame"
)

func main() {

	var exNumber = flag.Int("n", 1, "Sets which exercise is executed.")
	var timeOut = flag.Int("timeOut", 30, "Sets time limit for exercise in seconds.")
	var path = flag.String("pathEx1", "quizGame/problems.csv", "Set path to CSV file with questions for Exercise 1.")
	flag.Parse()

	if *exNumber == 1 {
		ex1.Ex1(*path, *timeOut)
	}

}
