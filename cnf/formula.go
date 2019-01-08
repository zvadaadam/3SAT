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





