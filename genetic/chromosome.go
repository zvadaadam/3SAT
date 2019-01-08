package genetic

import "math/rand"

type Chromosome struct {

	Genomes []bool

}

func NewChromosome(numGenomes int) *Chromosome {
	//chromosome := new(Chromosome)

	genomes := make([]bool, numGenomes)
	for i := 0; i < numGenomes; i++ {
		genomes[i] = randomBool()
	}

	return &Chromosome{Genomes: genomes}
	//return chromosome
}

func (ch *Chromosome) EvaluateFitness() int {
	// TODO
	return 0
}

func randomBool() bool {
	return (rand.Intn(2) == 0)
}