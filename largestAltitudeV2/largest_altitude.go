package main

import "fmt"

// 1732. Find the Highest Altitude

// Hint
// There is a biker going on a road trip. The road trip consists of n + 1 points at various altitudes.
//  The biker starts his trip on point 0 with altitude equal 0.

// You are given an integer array gain of length n where gain[i] is the net gain
//  in altitude between points i​​​​​​ and i + 1 for all (0 <= i < n).
//  Return the highest altitude of a point.
func largestAltitude(gain []int) int {
    if len(gain)==0{
        return 0
    }

    largest := 0
    altitude := 0

    for i, g := range gain{
        altitude = altitude + g
        fmt.Printf("count %d - %d", i, altitude)
        if altitude > largest{
            largest = altitude
        }
    }

    return largest
    
}