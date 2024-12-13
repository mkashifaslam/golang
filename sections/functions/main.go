package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 3, 7}
	fmt.Println("Normal", numbers)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("Doubled", doubled)
	fmt.Println("Tripled", tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	transformedMoreNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println("transformedNumbers", transformedNumbers)
	fmt.Println("transformedMoreNumbers", transformedMoreNumbers)

	// fmt.Println(numbers[:]) print full numbers array

	//transformed := transformNumbers(&numbers, func(number int) int {
	//	return number * 2
	//})
	//fmt.Println(transformed)

	doubledFunc := createTransformer(2)
	tripledFunc := createTransformer(3)
	doubledNumbers := transformNumbers(&numbers, doubledFunc)
	tripledNumbers := transformNumbers(&numbers, tripledFunc)

	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)

}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumbers []int

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
