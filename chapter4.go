// Linear search

package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time")

func main() {
    rand.Seed(time.Now().UnixNano())

    // Make and sort an array.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)
    arr := make_random_array(num_items, max)
    print_array(arr, 40)

    for {
        // Get the target as a string.
        var target_string string;
        fmt.Printf("Target: ")
        fmt.Scanln(&target_string)

        // If the target string is blank, break out of the loop.
        if len(target_string) == 0 { break }

        // Convert to int and find it.
        target, _ := strconv.Atoi(target_string)

        index, num_tests := linear_search(arr, target)
        if index < 0 || index >= len(arr) {
            fmt.Printf("Target %d not found, %d tests\n", target, num_tests)
        } else {
            fmt.Printf("arr[%d] = %d, %d tests\n", index, arr[index], num_tests)
        }
    }
}

// Make an array containing pseudorandom numbers in [0, max).
func make_random_array(num_items, max int) []int {
    arr := make([]int, num_items)
    for i := 0; i < num_items; i++ {
        arr[i] = rand.Intn(max)
    }
    return arr
}

// Print at most num_items items.
func print_array(arr []int, num_items int) {
    if len(arr) <= num_items {
        fmt.Println(arr)
    } else {
        fmt.Println(arr[:num_items])
    }
}

// Perform linear search.
// Return the target's location in the array and the number of tests.
// If the item is not found, return -1 and the number of tests.
func linear_search(arr []int, target int) (index, num_tests int) {
	for i, num := range arr {
		if num == target {
			return i, i + 1 // Return the index and the number of searches
		}
	}
	return -1, len(arr) // Return -1 if the target is not found and the number of searches
}

