package genetic

import (
	"SAT/cnf"
	"fmt"
)

func Solve(formula cnf.Formula, numGenerations int, populationSize int, tournamentSize int, mutationRate float32) ([]bool) {

	// Init population
	population := NewPopulation(populationSize, formula.NumVariables)

	for i := 0; i < numGenerations; i++ {

		parentChromosomeA, parentChromosomeB := population.SelectionTournament(formula, tournamentSize)

		childChromosome := parentChromosomeA.Crossover(&parentChromosomeB)

		childChromosome.Mutation(mutationRate)

		population.RemoveWeakest(formula)

		population.AppendChromosome(childChromosome)
		fmt.Printf("Fitness of Child = %d\n", childChromosome.EvaluateFitness(formula))
	}

	population.chromosomes = sortByFitness(population.chromosomes, formula)

	fmt.Print("____________________________________________\n")
	fmt.Printf("FINAL FITNESS: %d\n", population.chromosomes[0].EvaluateFitness(formula))

	var sumWeights = 0
	for _, val := range formula.Weights {
		sumWeights = sumWeights + val
	}
	fmt.Print(sumWeights)

	sumWeights = 0
	for i, weight := range formula.Weights {
		if population.chromosomes[0].Genomes[i] {
			sumWeights = sumWeights + weight
		}
	}
	fmt.Print(sumWeights)

	return population.chromosomes[0].Genomes
}
