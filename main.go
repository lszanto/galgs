package main

import "fmt"

func main() {
	// numbers to sort
	sortme := []int{239, 2923, 119, 440, 20, 38, 48, 2201, 2938, 20, 11, 44, 55, 66, 33, 203, 22}
	sortme2 := []int{239, 2923, 119, 440, 20, 38, 48, 2201, 2938, 20, 11, 44, 55, 66, 33, 203, 22}

	fmt.Println("********************************")
	fmt.Println("Sorted from: ")
	fmt.Println(sortme)
	iterations := sortArr(sortme)
	fmt.Println("To: ")
	fmt.Println(sortme)
	fmt.Printf("In %d iterations\n", iterations)

	fmt.Println("--------------------------------")
	fmt.Println("Sorted from: ")
	fmt.Println(sortme2)
	iterations = sortArr2(sortme2)
	fmt.Println("To: ")
	fmt.Println(sortme2)
	fmt.Printf("In %d iterations\n", iterations)
}

func sortArr(arr []int) int {
	// sorted
	sorted := false

	// count iterations
	iterations := 0

	// loop through
	for sorted == false {
		// set sorted to true
		sorted = true

		// loop through sortme
		for i := range arr {
			// add an iteration
			iterations++

			// grab a and b
			a := arr[i]

			// ensure not too big
			if (i + 1) < len(arr) {
				b := arr[i+1]

				// check if bigger
				if a > b {
					// switch
					arr[i] = b
					arr[i+1] = a

					// had to switch assume not true
					sorted = false
				}
			} else {
				continue
			}
		}
	}

	// return sorted
	return iterations
}

func sortArr2(arr []int) int {
	// sorted
	sorted := false

	// count iterations
	iterations := 0

	// current loop
	position := 1

	// loop through
	for sorted == false {
		// set sorted to true
		sorted = true

		// what we are comparing to
		smallest := arr[position-1]

		// loop through sortme
		for i := position; i < len(arr); i++ {
			// add an iteration
			iterations++

			// grab a
			a := arr[i]
			b := arr[position-1]

			// check if bigger
			if a < smallest {
				// switch
				arr[i] = b
				arr[position-1] = a

				// set new smallest
				smallest = a

				// had to switch assume not true
				sorted = false
			}
		}

		// move position on
		position++
	}

	// return sorted
	return iterations
}
