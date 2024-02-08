package calcEquation

//399. Evaluate Division
//Medium
//
//Topics
//Companies
//
//Hint
//You are given an array of variable pairs equations and an array of real numbers values,
//where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i].
//Each Ai or Bi is a string that represents a single variable.
//
//You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query
//where you must find the answer for Cj / Dj = ?.
//
//Return the answers to all queries. If a single answer cannot be determined, return -1.0.
//
//Note: The input is always valid. You may assume that evaluating the queries will
//not result in division by zero and that there is no contradiction.
//
//Note: The variables that do not occur in the list of equations are undefined,
//so the answer cannot be determined for them.

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i, eq := range equations {
		if _, ok := graph[eq[0]]; !ok {
			graph[eq[0]] = make(map[string]float64)
		}
		if _, ok := graph[eq[1]]; !ok {
			graph[eq[1]] = make(map[string]float64)
		}
		graph[eq[0]][eq[1]] = values[i]
		graph[eq[1]][eq[0]] = 1 / values[i]
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		res[i] = dfs(q[0], q[1], 1, graph, make(map[string]struct{}))
	}
	return res
}

func dfs(start, end string, cur float64, graph map[string]map[string]float64, visited map[string]struct{}) float64 {
	if _, ok := graph[start]; !ok {
		return -1
	}
	if start == end {
		return cur
	}
	visited[start] = struct{}{}
	for k, v := range graph[start] {
		if _, ok := visited[k]; !ok {
			if res := dfs(k, end, cur*v, graph, visited); res != -1 {
				return res
			}
		}
	}
	return -1
}
