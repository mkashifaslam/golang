package main

import (
	"fmt"
	priceslist "github.com/mkashifaslam/golang/price-calculator/prices-list"
	taxcalc "github.com/mkashifaslam/golang/price-calculator/tax-calculator"
)

func main() {
	prices, err := priceslist.GetPricesFromFile("sections/projects/price-calculator/prices.txt") //[]int{15, 25, 35}
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRates := []int{0, 10, 20}
	result := taxcalc.CalculateTaxes(prices, taxRates)
	err = taxcalc.Save(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully saved result!")
}
