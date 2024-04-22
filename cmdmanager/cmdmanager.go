package cmdmanager

import "fmt"

// CMDManager user to only connect methods
type CMDManager struct {
}

func (fm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Print enter your prices. Confirm every prices with enter")
	var prices []string
	for {
		var price string
		fmt.Printf("Price:")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (fm CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
