package prices

import (
	"fmt"
	"go_calculator/conversion"
	"go_calculator/iomanager"
)

// struct for managing the prices
type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

// NewTaxIncludedPriceJob func to construct struct, usually called New
// *constructor function
// it returns the tax included prices pointer address//
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

// method
func (job *TaxIncludedPriceJob) Process() error {
	//calling LoadData to make sure we have data we need to load
	err := job.LoadData() //loads the data from the document to be read before making any changes

	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	//write in json and uses the job struct to save it on result.txt

	//let dynamically generate the file names to avoid collisions in the result
	//that is used to differentiate each result
	return job.IOManager.WriteResult(job)

}

// Loader function to load the prices from a prices.txt
func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Failed to read prices.txt")
		fmt.Println(err)
		return err
	}
	//will convert all prices to float64
	price, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Println("FAILED to convert")

		return err
	}

	job.InputPrices = price

	return nil

}
