package main

import (
	"fmt"
	"github.com/mkashifaslam/golang/projects/price-calculator-max/cmdmanager"
	"github.com/mkashifaslam/golang/projects/price-calculator-max/prices"
)

var inputFile = "prices.txt"

var getOutputFile = func(taxRate float64) string {
	return fmt.Sprintf("result_%.0f.json", taxRate*100)
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		//fm := filemanager.New(inputFile, getOutputFile(taxRate))
		cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		priceJob.Process()
	}
}
