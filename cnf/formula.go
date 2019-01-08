package cnf

type Formula struct {
	Clauses [][]int
	NumClauses int
	NumVariables int
}

func (f *Formula) AddClauses(clauses [][]int, numVariables int, numClauses int) {

	f.Clauses = clauses
	f.NumClauses = numClauses
	f.NumVariables = numVariables
}


func (f *Formula) Satisfied(variables []bool) (bool, int) {

	var sumTrues = 0

	for i := 0; i < f.NumClauses; i++ {

		var isClauseSatisfied = false
		var variablesInClause = len(f.Clauses[i])
		for j := 0; j < variablesInClause; j++ {
			literal := f.Clauses[i][j]
			variable := variables[abs(literal) - 1]

			if evalLiteral(literal, variable) {
				sumTrues = sumTrues + 1
				isClauseSatisfied = true
			}
		}

		if !isClauseSatisfied {
			return false, 0
		}
	}

	return true, sumTrues
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