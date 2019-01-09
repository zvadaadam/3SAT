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

			weight := f.Weights[literalIndex]
			variable := variables[literalIndex]

			if evalLiteral(literal, variable) {
				sumWeights = sumWeights + weight
				isClauseSatisfied = true
			}
		}

		if !isClauseSatisfied {
			sumNotSatsfiedClauses = sumNotSatsfiedClauses - 1
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