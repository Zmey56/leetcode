package minReorder

//There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two
//different cities (this network form a tree). Last year, The ministry of transport decided to orient the roads
//in one direction because they are too narrow.
//
//Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.
//
//This year, there will be a big event in the capital (city 0), and many people want to travel to this city.
//
//Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number
//of edges changed.
//
//It's guaranteed that each city can reach city 0 after reorder.

func minReorder(n int, connections [][]int) int {
	forward := make(map[int][]int)
	reverse := make(map[int][]int)
	for _, conn := range connections {
		forward[conn[0]] = append(forward[conn[0]], conn[1])
		reverse[conn[1]] = append(reverse[conn[1]], conn[0])
	}

	visited := make([]bool, n)
	var dfs func(int) int
	dfs = func(node int) int {
		visited[node] = true
		count := 0
		for _, next := range forward[node] {
			if !visited[next] {
				count += 1 + dfs(next) // Need to reverse this edge
			}
		}
		for _, prev := range reverse[node] {
			if !visited[prev] {
				count += dfs(prev) // No need to reverse this edge
			}
		}
		return count
	}

	return dfs(0)
}
