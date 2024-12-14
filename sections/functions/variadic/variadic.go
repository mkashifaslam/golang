package main

import "fmt"

func main() {
	numbers := []int{1, 20, 5}
	//sum := sumup(1, 20, 13, -4, 11)
	sum := sumup(numbers...)
	fmt.Println("Sum of these numbers is ", sum)
}

func sumup(numbers ...int) int {
	sum := 0

	fmt.Println("Total numbers are ", len(numbers))

	for _, val := range numbers {
		sum += val
	}

	return sum
}
