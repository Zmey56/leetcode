package main

//46. Permutations

// Given an array nums of distinct integers, return all the possible permutations.
//  You can return the answer in any order.
func permute(nums []int) [][]int {
    result := [][]int{}

    var backtrack func(path []int, used []bool) 

    backtrack = func(path []int, used []bool) {
        if len(path)==len(nums){
            temp := make([]int, len(path))
            copy(temp, path)
            result = append(result, temp)
            return 
        }

        for i := 0; i< len(nums); i++{
            if used[i]{
                continue
            }

            used[i] = true
            path = append(path, nums[i])

            backtrack(path, used)

            path = path[:len(path)-1]
            used[i] = false
        }
    }

    backtrack([]int{}, make([]bool, len(nums)))

    return result
    
}