package slices

import "fmt"

func SliceList() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices))

	prices[1] = 10.32
	prices = append(prices, 32.11)
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices))
}
