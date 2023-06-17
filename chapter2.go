// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // Get the number of items and maximum item value.
	var num_items = 5
	var max = 10

    // Make and display the unsorted array.
    arr := make_random_array(num_items, max)
    print_array(arr, 40)
    fmt.Println()

	// Sort and display the result.
	quicksort(arr)
	print_array(arr, 40)

	// Verify that it's sorted.
	check_sorted(arr)
}

func check_sorted(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}

	fmt.Println("The array is sorted")
}

func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	quicksort(arr[:i])
	quicksort(arr[i+1:])
}

func print_array(arr []int, length int) {
	for i := 0; i < len(arr) && i < length; i++ {
		fmt.Println(arr[i])
	}
}

func make_random_array(num_items, max int) []int {
	arr := make([]int, num_items)
	for i := 0; i < num_items; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}
