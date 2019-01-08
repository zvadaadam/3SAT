package genetic

import (
	"SAT/cnf"
	"fmt"
)

func Solve(formula cnf.Formula, numGenerations int, populationSize int, tournamentSize int, mutationRate float32) {

	// Init population
	population := NewPopulation(populationSize, formula.NumVariables)
	fmt.Print(population)

	//for i := 0; i < numGenerations; i++ {
	//	// selection
	//	// crossover
	//	// muation
	//	// elimination
	//
	//}

}
