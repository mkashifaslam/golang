package arrays

import "fmt"

func Array() {
	prices := [4]float64{11.32, 62.32, 83.12, 40.22}
	fmt.Println(prices)
	var books = [4]string{"The book"}

	books[2] = "The carpet"

	fmt.Println(prices[2])

	fmt.Println(books)

	featuresPrices := prices[1:]
	fmt.Println(featuresPrices)

	highlightedPrices := featuresPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

}

func ListUnpack() {
	prices := []float64{11.32, 62.32, 83.12, 40.22}
	fmt.Println(prices)

	discountPrices := []float64{5.31, 2.42, 8.1}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
