package tax_calculator

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaxPrices struct {
	Tax    int   `json:"tax"`
	Prices []int `json:"prices"`
}

func (tp *TaxPrices) GetFormattedStr() string {
	return fmt.Sprintf("%v %v\n", tp.Tax, tp.Prices)
}

func Save(taxedPrices []TaxPrices) error {
	fileName := "taxed_prices.json"

	jsonValue, err := json.Marshal(taxedPrices)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonValue, 0644)
}

func taxCalculator(prices []int, tax int) (result TaxPrices) {
	var taxedPrices []int
	for _, price := range prices {
		taxedPrices = append(taxedPrices, price+price*tax/100)
	}

	return TaxPrices{
		Tax:    tax,
		Prices: taxedPrices,
	}
}

func CalculateTaxes(prices []int, taxes []int) (taxedPrices []TaxPrices) {
	for _, tax := range taxes {
		taxedPrices = append(taxedPrices, taxCalculator(prices, tax))
	}

	return taxedPrices
}
