package prices

import (
	"bufio"
	"fmt"
	"go_calculator/conversion"
	"os"
)

// struct for managing the prices
type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// method
func (job *TaxIncludedPriceJob) Process() {
	//calling LoadData to make sure we have data we need to load
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.f", taxIncludedPrice)
	}

	fmt.Println(result)
}

// Loader function to load the prices from a prices.txt
func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	//opening file
	if err != nil {
		fmt.Println("Could not  open:", err)
		return
	}
	//reading line by line content with scanner
	scanner := bufio.NewScanner(file) //scanner line by line

	var lines []string //slice for each line
	//for loop to append each line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	//store error in already initialized err variable
	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading was not possible open:", err)
		file.Close()
		return
	}

	//converting prices to float64

	//will convert all prices to float64
	price, err := conversion.StringToFloats(lines)

	if err != nil {
		fmt.Println("FAILED to convert")
		file.Close()
		return
	}

	//for lineIndex, line := range lines {
	//	floatPrice, err := strconv.ParseFloat(line, 64)
	//
	//	if err != nil {
	//		fmt.Println("converting was not possible open:", err)
	//		file.Close()
	//		return
	//	}
	//
	//	price[lineIndex] = floatPrice
	//}
	job.InputPrices = price
	file.Close()
}

// NewTaxIncludedPriceJob func to construct struct, usually called New
// *constructor function
// it returns the tax included prices pointer address//
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
