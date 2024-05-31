package isToeplitzMatrix

import "sync"

//Given an m x n matrix, return true if the matrix is Toeplitz. Otherwise, return false.
//
//A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same elements.

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true

}

func isToeplitzMatrixVer2(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])

	// Create a channel to receive results from goroutines.
	results := make(chan bool, m+n-1)

	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Check each diagonal in a separate goroutine.
	for d := 0; d < m+n-1; d++ {
		wg.Add(1)
		go func(d int) {
			defer wg.Done()

			var value int
			if d < m {
				value = matrix[d][0]
			} else {
				value = matrix[0][d-m+1]
			}

			for i := 0; i <= d && i < m && d-i < n; i++ {
				if matrix[i][d-i] != value {
					results <- false
					return
				}
			}

			results <- true
		}(d)
	}

	// Close the results channel after all goroutines have finished.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Check the results from the channel.
	for result := range results {
		if !result {
			return false
		}
	}

	return true
}
