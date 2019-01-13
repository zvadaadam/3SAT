package genetic

import (
	"SAT/cnf"
	"fmt"
)

func Solve(formula cnf.Formula, numGenerations int, populationSize int, elitismSize int, randomSize int, tournamentSize int, mutationRate float64) ([]bool) {

	// Init population
	population := NewRandomPopulation(populationSize, formula.NumVariables)
	bestIndividual := population.chromosomes[0]
	for i := 0; i < numGenerations; i++ {

		newPopulation := NewRandomPopulation(randomSize, formula.NumVariables)

		fittests := population.FittestIndividuals(elitismSize, formula)

		for _, individual := range fittests {
			newPopulation.AppendChromosome(individual)
		}

		numChilds := populationSize - randomSize - elitismSize
		if numChilds <= 0 {
			fmt.Print("No more breeding, no more life...")
			return []bool{}
		}

		for i := 0; i < numChilds; i++ {
			parentChromosomeA, parentChromosomeB := population.SelectionTournament(formula, tournamentSize)
			//childChromosome := parentChromosomeA.Crossover(&parentChromosomeB)
			childChromosome := parentChromosomeA.CrossoverTwoPoint(&parentChromosomeB)
			childChromosome.Mutation(mutationRate)
			newPopulation.AppendChromosome(childChromosome)

			population = newPopulation
		}

		bestIndividualPopulation := population.FittestIndividuals(1, formula)[0]
		if bestIndividual.EvaluateFitness(formula) < bestIndividualPopulation.EvaluateFitness(formula) {
			bestIndividual = bestIndividualPopulation
		}
		//fmt.Printf("Fitness of Child = %d\n", population.FittestIndividuals(1, formula)[0].EvaluateFitness(formula))
	}

	//population.chromosomes = sortByFitness(population.chromosomes, formula)

	//fmt.Print("____________________________________________\n")
	//fmt.Printf("FINAL FITNESS: %d\n", population.chromosomes[0].EvaluateFitness(formula))

	//return population.chromosomes[0].Genomes

	fmt.Printf("%v", bestIndividual.EvaluateFitness(formula))

	return bestIndividual.Genomes
}
