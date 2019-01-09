package cnf

type Formula struct {
	Clauses [][]int
	Weights []int
	NumClauses int
	NumVariables int
}

func (f *Formula) AddClauses(clauses [][]int, weights []int, numVariables int, numClauses int) {
	f.Clauses = clauses
	f.Weights = weights
	f.NumClauses = numClauses
	f.NumVariables = numVariables
}


func (f *Formula) Satisfied(variables []bool) (int, int) {

	var sumWeights = 0
	var sumNotSatsfiedClauses = 0

	for i := 0; i < f.NumClauses; i++ {

		var isClauseSatisfied = false
		var variablesInClause = len(f.Clauses[i])
		for j := 0; j < variablesInClause; j++ {
			literal := f.Clauses[i][j]
			literalIndex := abs(literal) - 1

			variable := variables[literalIndex]
			if evalLiteral(literal, variable) {
				isClauseSatisfied = true
			}
		}

		if !isClauseSatisfied {
			sumNotSatsfiedClauses = sumNotSatsfiedClauses - 1
		}
	}

	for i, weight := range f.Weights {
		if variables[i] {
			sumWeights = sumWeights + weight
		}
	}

	return sumNotSatsfiedClauses, sumWeights
}

func evalLiteral(literal int, variable bool) bool {
	if literal < 0 {
		return !variable
	}
	return variable
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}