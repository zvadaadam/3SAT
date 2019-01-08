package genetic


type Population struct {

	chromosomes []Chromosome
}

func NewPopulation(populationSize int, numGenomes int) *Population {
	population := new(Population)

	chromosomes := make([]Chromosome, 0)
	for i := 0; i < populationSize; i++ {
		chromosome := NewChromosome(numGenomes)
		chromosomes = append(chromosomes, *chromosome)
	}

	population.chromosomes = chromosomes

	return population
}



