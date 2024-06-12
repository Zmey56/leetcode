package validMountainArray

//Given an array of integers arr, return true if and only if it is a valid mountain array.
//
//Recall that arr is a mountain array if and only if:
//
//arr.length >= 3
//There exists some i with 0 < i < arr.length - 1 such that:
//arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	mountain := true
	high := true
	top := false
	prev := arr[0]

	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] == prev {
			mountain = false
			break
		}

		if arr[i] > prev && high {
			prev = arr[i]
		}

		if arr[i] < prev && high && i != 1 {
			high = false
			top = true
			prev = arr[i]
		} else if arr[i] < prev && high && i == 1 {
			mountain = false
		}

		if arr[i] < prev && !high {
			prev = arr[i]
		}

		if arr[i] > prev && !high {
			mountain = false
			break
		}

	}

	return mountain && top
}

func validMountainArrayVer2(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	n := len(arr)
	asc := make(chan bool)
	desc := make(chan bool)

	go func() {
		i := 0
		for i < n-1 && arr[i] < arr[i+1] {
			i++
		}
		if i == n-1 || i == 0 {
			asc <- false
		} else {
			asc <- true
		}
		close(asc)
	}()

	go func() {
		i := 0
		for i < n-1 && arr[i] < arr[i+1] {
			i++
		}
		if i == n-1 || i == 0 {
			desc <- false
		} else {
			j := i
			for j < n-1 && arr[j] > arr[j+1] {
				j++
			}
			if j == n-1 {
				desc <- true
			} else {
				desc <- false
			}
		}
		close(desc)
	}()

	return <-asc && <-desc
}
