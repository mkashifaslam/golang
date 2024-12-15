package main

import (
	"fmt"
	"github.com/mkashifaslam/golang/projects/price-calculator-max/filemanager"
	"github.com/mkashifaslam/golang/projects/price-calculator-max/prices"
)

var inputFile = "prices.txt"

var getOutputFile = func(taxRate float64) string {
	return fmt.Sprintf("result_%.0f.json", taxRate*100)
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New(inputFile, getOutputFile(taxRate))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		//if err != nil {
		//	fmt.Println("could not process job")
		//	fmt.Println(err)
		//}
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			fmt.Println(err)
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
