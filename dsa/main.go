package main

import "fmt"

func binarySearch(arr []int, target int) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := low + (high-low) / 2
        if arr[mid] == target {
            return mid
        }
        if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1
}

func main() {
    arr := []int{2, 3, 4, 10, 40}
    target := 10
    result := binarySearch(arr, target)
    if result != -1 {
        fmt.Println("Element found at index", result)
    } else {
        fmt.Println("Element not found")
    }
}
