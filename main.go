package main

import (
	"flag"

	"gophercises.gonzalogarcia.de/ex1"
)

func main() {

	var exNumber = flag.Int("n", 1, "Sets which exercise is executed.")
	var timeOut = flag.Int("timeOut", 30, "Sets time limit for exercise in seconds.")
	var path = flag.String("pathEx1", "ex1/problems.csv", "Set path to CSV file with questions for Exercise 1.")
	flag.Parse()

	if *exNumber == 1 {
		ex1.Ex1(*path, *timeOut)
	}

}
