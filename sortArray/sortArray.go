package main

// 912. Sort an Array

// Given an array of integers nums, sort the array in ascending order and return it.

// You must solve the problem without using any built-in functions in O(nlog(n)) time
//  complexity and with the smallest space complexity possible.
func sortArray(nums []int) []int {
    // Базовый случай: массив из 1 или 0 элементов уже отсортирован
    if len(nums) <= 1 {
        return nums
    }
    
    // Делим массив ровно пополам
    mid := len(nums) / 2
    
    // Рекурсивно сортируем левую и правую части
    left := sortArray(nums[:mid])
    right := sortArray(nums[mid:])
    
    // Сливаем две отсортированные части в одну
    return merge(left, right)
}

// Функция для склеивания двух отсортированных массивов
func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0
    
    // Сравниваем элементы из обеих половин и берем меньший
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    
    // Докидываем остатки, если одна из половин закончилась раньше
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    
    return result
}