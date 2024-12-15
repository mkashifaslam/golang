package prices

import (
	"encoding/json"
	"fmt"
	"github.com/mkashifaslam/golang/projects/price-calculator-max/conversion"
	"github.com/mkashifaslam/golang/projects/price-calculator-max/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = taxIncludedPrice
	}

	job.TaxIncludedPrices = result

	err := job.IOManager.WriteResult(job)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (job *TaxIncludedPriceJob) Print() {
	jobJsonEncoding, _ := json.Marshal(job)
	fmt.Println(string(jobJsonEncoding))
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
