package nearestExit

import "container/list"

// You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+').
// You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.
//
// In one step, you can move one cell up, down, left, or right. You cannot step into a cell with a wall, and you cannot
// step outside the maze. Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell
// that is at the border of the maze. The entrance does not count as an exit.
//
// Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.
func nearestExit(maze [][]byte, entrance []int) int {
	// Create a queue to store the coordinates of the cells
	queue := make([][]int, 0)
	queue = append(queue, entrance)

	// Create a visited set to store the visited cells
	visited := make(map[[2]int]bool)
	visited[[2]int{entrance[0], entrance[1]}] = true

	// Create a directions array to store the directions
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// Create a variable to store the number of steps
	steps := 0

	// Iterate over the queue
	for len(queue) > 0 {
		// Increment the number of steps
		steps++

		// Create a variable to store the size of the queue
		size := len(queue)

		// Iterate over the queue
		for i := 0; i < size; i++ {
			// Dequeue the cell
			cell := queue[0]
			queue = queue[1:]

			// Iterate over the directions
			for _, direction := range directions {
				// Calculate the new row and column
				newRow := cell[0] + direction[0]
				newCol := cell[1] + direction[1]

				// Check if the new row and column are within the bounds of the maze
				if newRow >= 0 && newRow < len(maze) && newCol >= 0 && newCol < len(maze[0]) {
					// Check if the cell is empty and not visited
					if maze[newRow][newCol] == '.' && !visited[[2]int{newRow, newCol}] {
						// Check if the cell is at the border of the maze
						if newRow == 0 || newRow == len(maze)-1 || newCol == 0 || newCol == len(maze[0])-1 {
							return steps
						}

						// Mark the cell as visited
						visited[[2]int{newRow, newCol}] = true

						// Enqueue the cell
						queue = append(queue, []int{newRow, newCol})
					}
				}
			}
		}
	}

	// Return -1 if no exit is found
	return -1
}

func nearestExitVer2(maze [][]byte, entrance []int) int {
	queue := list.New()
	queue.PushBack(entrance)

	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}
	visited[entrance[0]][entrance[1]] = true

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	steps := 0

	for queue.Len() > 0 {
		steps++
		size := queue.Len()

		for i := 0; i < size; i++ {
			cell := queue.Remove(queue.Front()).([]int)

			for _, direction := range directions {
				newRow := cell[0] + direction[0]
				newCol := cell[1] + direction[1]

				if newRow >= 0 && newRow < len(maze) && newCol >= 0 && newCol < len(maze[0]) {
					if maze[newRow][newCol] == '.' && !visited[newRow][newCol] {
						if newRow == 0 || newRow == len(maze)-1 || newCol == 0 || newCol == len(maze[0])-1 {
							return steps
						}

						visited[newRow][newCol] = true
						queue.PushBack([]int{newRow, newCol})
					}
				}
			}
		}
	}

	return -1
}
