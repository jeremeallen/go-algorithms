/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import (
    "fmt"
    "math/rand"
    "time")

func make_random_array(num_items, max int) []int {
    arr := make([]int, num_items)
    for i := 0; i < num_items; i++ {
        arr[i] = rand.Intn(max)
    }
    return arr
}

func print_array(arr []int, num_items int) {
    if len(arr) <= num_items {
        fmt.Println(arr)
    } else {
        fmt.Println(arr[:num_items])
    }
}

func check_sorted(arr []int) {
    for i := 1; i < len(arr); i++ {
        if arr[i - 1] > arr[i] {
            fmt.Println("The array is NOT sorted!")
            return
        }
    }

    fmt.Println("The array is sorted")
}

func bubble_sort(arr []int) {
	fmt.Println("Bubble Sort")
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
    bubble_sort(arr)
    print_array(arr, 40)

    // Verify that it's sorted.
    check_sorted(arr)
}