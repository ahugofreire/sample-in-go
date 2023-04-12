package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

func main() {

}

// map
func funcMap() {
	nums := []int{1, 2, 3, 4, 5}

	result := funk.Map(nums, func(n int) int {
		return n * 2
	})

	fmt.Println(result) // [2 4 6 8 10]
}

func funcFilter() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filtered := funk.FilterInt(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(filtered) // Output: [2 4 6 8 10]
}
