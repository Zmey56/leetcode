package calPoints

import (
	"strconv"
)

//You are keeping the scores for a baseball game with strange rules. At the beginning of the game, you start with an empty record.
//
//You are given a list of strings operations, where operations[i] is the ith operation you must apply to the record and is one of the following:
//
//An integer x.
//Record a new score of x.
//'+'.
//Record a new score that is the sum of the previous two scores.
//'D'.
//Record a new score that is the double of the previous score.
//'C'.
//Invalidate the previous score, removing it from the record.
//Return the sum of all the scores on the record after applying all the operations.
//
//The test cases are generated such that the answer and all intermediate calculations fit in a 32-bit integer and that all operations are valid.

func calPoints(operations []string) int {
	result := []int{}
	for _, op := range operations {
		switch op {
		case "+":
			result = append(result, result[len(result)-1]+result[len(result)-2])
		case "D":
			result = append(result, result[len(result)-1]*2)
		case "C":
			result = result[:len(result)-1]
		default:
			val, _ := strconv.Atoi(op)
			result = append(result, val)
		}
	}

	sum := 0

	for _, v := range result {
		sum += v
	}

	return sum

}
