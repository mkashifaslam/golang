package cmdmanager

import "fmt"

type CmdManager struct {
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter prices, confirm each price with enter")

	var prices []string
	for {
		var price string
		fmt.Print("Enter price ")
		fmt.Scan(&price)

		if price == "" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CmdManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CmdManager {
	return CmdManager{}
}
