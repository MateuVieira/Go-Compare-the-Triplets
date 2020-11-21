package main

import "fmt"

// Complete the compareTriplets function below.
func compareTriplets(a []int, b []int) []int {
	result := []int{0, 0}

	for index, value := range a {
		if value > b[index] {
			result[0]++
		}

		if value < b[index] {
			result[1]++
		}
	}

	return result
}

func main() {
	a := []int{5, 6, 7}

	b := []int{3, 6, 10}

	result := compareTriplets(a, b)

	fmt.Println("Result: Alice: ", result[0], "  Bob: ", result[0])

}
