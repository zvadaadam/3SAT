package main

import (
	"SAT/cnf"
	"SAT/genetic"
	"fmt"
	"os"
)

func main() {

	fmt.Print("HELLO WORLD\n")

	// TODO: change to run arugument
	filePath := "/Users/adamzvada/go/src/SAT/input/dimacs_2.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	formula, err := cnf.Parse(file)

	if err != nil {
		fmt.Print(err)
		return
	}

	var numGenerations = 1000
	var populationSize = 50
	var tournamentSize = 10
	var mutationRate float32 = 0.8

	solution := genetic.Solve(formula, numGenerations, populationSize, tournamentSize, mutationRate)

	fmt.Print(solution)

	return
}