// Counting sort

package main

import (
    "fmt"
    "math/rand"
    "time")

type Customer struct {
    id              string
    num_purchases   int
}

func main() {
    rand.Seed(time.Now().UnixNano())

    // Get the number of items and maximum item value.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted array.
    arr := make_random_array(num_items, max)
    print_array(arr, 40)
    fmt.Println()

    // Sort and display the result.
    sorted := counting_sort(arr, max)
    print_array(sorted, 40)

    // Verify that it's sorted.
    check_sorted(sorted)
}


// Make an array containing pseudorandom
// customers with num_purchases in [0, max).
// The ith customer has id C<i> as in C103.
func make_random_array(num_items, max int) []Customer {
    arr := make([]Customer, num_items)
    for i := 0; i < num_items; i++ {
        id := fmt.Sprintf("C%d", i)
        num_purchases := rand.Intn(max)
        arr[i] = Customer { id, num_purchases }
    }
    return arr
}

// Print at most num_items items.
func print_array(arr []Customer, num_items int) {
    if len(arr) <= num_items {
        fmt.Println(arr)
    } else {
        fmt.Println(arr[:num_items])
    }
}

// Verify that the array is sorted.
func check_sorted(arr []Customer) {
    for i := 1; i < len(arr); i++ {
        if arr[i - 1].num_purchases > arr[i].num_purchases {
            fmt.Println("The array is NOT sorted!")
            return
        }
    }

    fmt.Println("The array is sorted")
}

// Sort an array containing items in [0, max).
func counting_sort(arr []Customer, max int) ([]Customer) {
    num_items := len(arr)

    // Make an array to hold counts.
    counts := make([]int, max)

    // Count the values.
    for i := 0; i < num_items; i++ {
        counts[arr[i].num_purchases]++
    }

    // Convert counts into counts of values <=.
    for i := 1; i < max; i++ {
        counts[i] += counts[i - 1]
    }

    // Count out the values.
    result := make([]Customer, num_items)
    for i := num_items - 1; i >= 0; i-- {
        // Move item i into position.
        num := arr[i].num_purchases
        result[counts[num] - 1] = arr[i]
        counts[num] -= 1
    }

    return result
}

