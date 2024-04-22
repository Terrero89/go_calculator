package main

import (
	"fmt"
	"go_calculator/filemanager"
	"go_calculator/prices"
)

func main() {

	var taxRates = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Printf("Could not process")
			fmt.Println(err)
			return
		}
	}
}
