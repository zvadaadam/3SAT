package cnf

import (
	"bytes"
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Parse(r io.Reader) (Formula, error) {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var formula = Formula{}

	for scanner.Scan() {
		token := scanner.Bytes()

		if len(token) == 0 {
			continue
		}

		switch string(token[0]) {
		case "c":
			continue
		case "p":
			fields := bytes.Fields(token)

			var (
				numLiterals = 0
				numClauses = 0
				clauses = [][]int{}
				weights = []int{}
				err error = nil
			)

			if numLiterals, err = strconv.Atoi(string(fields[2])); err != nil {
				return Formula{}, fmt.Errorf("Error on Number of Literals: %q", fields[2])
			}

			if numClauses, err = strconv.Atoi(string(fields[3])); err != nil {
				return Formula{}, fmt.Errorf("Error on Number of Clauses: %q", fields[3])
			}

			if clauses, weights, err = parseForumula(scanner, numClauses, numLiterals); err != nil {
				return Formula{}, err
			}

			formula.AddClauses(clauses, weights, numLiterals, numClauses)
		default:
			return Formula{}, fmt.Errorf("invalid start of line character: %q", string(token[0]))

		}
	}

	return formula, nil
}

func parseForumula(scanner *bufio.Scanner, numClauses int, numLiterals int) ([][]int, []int, error) {

	clauses := make([][]int, 0)
	weights := make([]int, 0)

	scanner.Scan()
	token := scanner.Bytes()
	fields := bytes.Fields(token)

	if string(fields[0]) != "w" {
		return clauses, weights, fmt.Errorf("Error - Missing weights row!")
	}

	for _, field := range fields[1:] {
		var err error
		var weight = 0
		if weight, err = strconv.Atoi(string(field)); err != nil {
			return clauses, weights, fmt.Errorf("Error expecting integer weight: %q", weight)
		}

		weights = append(weights, weight)
	}

	if len(weights) != numLiterals {
		return clauses, weights, fmt.Errorf("Number of Weights should be to number of literals")
	}

	for scanner.Scan() {
		var err error
		token := scanner.Bytes()
		fields := bytes.Fields(token)

		literals := make([]int, 0)
		for _, field := range fields {
			var val = 0
			if val, err = strconv.Atoi(string(field)); err != nil {
				return clauses, weights, fmt.Errorf("Error expecting integer literal: %q", val)
			}

			if val > numClauses {
				return clauses, weights, fmt.Errorf("Literal identifier exceecds max value: %q", val)
			}

			if val != 0 {
				literals = append(literals, val)
			}
		}

		clauses = append(clauses, literals)
	}

	return clauses, weights, nil
}