package genetic

import (
	"SAT/cnf"
	"math/rand"
	"sort"
)

type Population struct {

	chromosomes []Chromosome
}

func NewRandomPopulation(populationSize int, numGenomes int) *Population {
	population := new(Population)

	chromosomes := make([]Chromosome, 0)
	for i := 0; i < populationSize; i++ {
		chromosome := NewChromosome(numGenomes)
		chromosomes = append(chromosomes, *chromosome)
	}

	population.chromosomes = chromosomes

	return population
}

func (p *Population) AppendPopolation(chromosome Chromosome) {



}

func (p *Population) SelectionTournament(formula cnf.Formula, tournamentSize int) (Chromosome, Chromosome) {

	tournament := make([]Chromosome, 0)
	for i := 0; i < tournamentSize; i++ {
		rndIndex := rand.Intn(len(p.chromosomes))
		tournament = append(tournament, p.chromosomes[rndIndex])
	}

	tournament = sortByFitness(tournament, formula)

	return p.chromosomes[0], p.chromosomes[1]
}

func (p *Population) SelectionRoulette(formula cnf.Formula) Chromosome {

	var sumFitness = 0
	for _, chromosome := range p.chromosomes {
		sumFitness = sumFitness + chromosome.EvaluateFitness(formula)
	}

	var cursor = rand.Float32()*float32(sumFitness)
	for _, chromosome := range p.chromosomes {
		fitness := chromosome.EvaluateFitness(formula)
		cursor = cursor - float32(fitness)
		if cursor < 0 {
			return chromosome
		}
	}

	return p.chromosomes[len(p.chromosomes) - 1]
}

func (p *Population) FittestIndividuals(numIndividuals int, formula cnf.Formula) []Chromosome {
	p.chromosomes = sortByFitness(p.chromosomes, formula)

	return p.chromosomes[0:numIndividuals]
}

func (p *Population) RemoveWeakest(formula cnf.Formula) {
	p.chromosomes = sortByFitness(p.chromosomes, formula)[:len(p.chromosomes)-1]
}

func (p *Population) AppendChromosome(chromosome Chromosome) {
	p.chromosomes = append(p.chromosomes, chromosome)
}

func sortByFitness(chromosomes []Chromosome, formula cnf.Formula) []Chromosome {

	sort.Slice(chromosomes, func(i, j int) bool {
		leftChromosomeFitness := chromosomes[i].EvaluateFitness(formula)
		rightChromosomeFitness  := chromosomes[j].EvaluateFitness(formula)

		return leftChromosomeFitness > rightChromosomeFitness
	})

	return chromosomes
}

