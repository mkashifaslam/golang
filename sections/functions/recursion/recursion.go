package recursion

import (
	"fmt"
	"math/big"
)

func Recursion() {
	//number := 100
	//fact := factorial1(number)
	number := int64(10) // Change this for different numbers
	fact := factorial(number)
	fmt.Printf("Factorial of %d is %d\n", number, fact)
}

func factorial(number int64) int64 {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func factorialBig(number int64) *big.Int {
	// Base case: 0! = 1
	if number == 0 {
		return big.NewInt(1)
	}

	// Recursive case: number * factorial(number-1)
	n := big.NewInt(number)
	return n.Mul(n, factorialBig(number-1))
}

// 9682165104862298112
// 15188249005818642432
