package genetic

import (
	"SAT/cnf"
	"math/rand"
)

type Chromosome struct {

	Genomes []bool

}

func NewChromosome(numGenomes int) *Chromosome {

	genomes := make([]bool, numGenomes)
	for i := 0; i < numGenomes; i++ {
		genomes[i] = randomBool()
	}

	return &Chromosome{Genomes: genomes}
}

func (ch *Chromosome) EvaluateFitness(formula cnf.Formula) int {

	sumOfNotStatisfied, sumTrues := formula.Satisfied(ch.Genomes)

	if sumOfNotStatisfied < 0 {
		return sumOfNotStatisfied
	}

	return sumTrues
}

func (ch *Chromosome) Crossover(chromosomeB *Chromosome) Chromosome {
	numGenomes := len(ch.Genomes)
	pivot := rand.Intn(numGenomes)

	child := make([]bool, numGenomes)
	for i := 0; i < numGenomes; i++ {
		if i < pivot {
			child[i] = ch.Genomes[i]
		} else {
			child[i] = chromosomeB.Genomes[i]
		}
	}

	return Chromosome{Genomes: child}
}

func (ch * Chromosome) Mutation(mutationRate float32) {

	if rand.Intn(100) <= int(mutationRate*100) {
		rndIndex := rand.Intn(len(ch.Genomes))
		ch.Genomes[rndIndex] = !ch.Genomes[rndIndex]
	}
}

func randomBool() bool {
	return (rand.Intn(2) == 0)
}