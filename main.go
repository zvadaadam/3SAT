package main

import (
	"SAT/cnf"
	"SAT/genetic"
	"flag"
	"fmt"
	"os"
)

func main() {

	inputArg := flag.String("input", "", "absolute path to the input file")
	generationArg := flag.Int("generation", 500, "Number of generations.")
	populationArg := flag.Int("population", 50, "The population size.")
	randomArg := flag.Int("random", 10, "Number of random individuals in population.")
	elitismArg := flag.Int("elitism", 5, "Number of the fittest individuals passed to the next generations.")
	tournamentArg := flag.Int("tournament", 10, "Number of contestants in selection tournament.")
	mutationArg := flag.Float64("mutation", 0.1, "The mutation rate.")

	flag.Parse()

	//input := string("/Users/adamzvada/go/src/SAT/input/hard")
	//file, err := os.Open(input)
	file, err := os.Open(*inputArg)
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

	var numGenerations = *generationArg
	var populationSize = *populationArg
	var tournamentSize = *tournamentArg
	var elitismSize = *elitismArg
	var randomSize = *randomArg
	var mutationRate = *mutationArg

	solution := genetic.Solve(formula, numGenerations, populationSize, randomSize, elitismSize, tournamentSize, mutationRate)

	fmt.Print(solution)

	return
}