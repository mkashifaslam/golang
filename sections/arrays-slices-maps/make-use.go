package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func MakeUse() {
	//userNames := []string{} // got index out of range error
	userNames := make([]string, 2, 5)
	userNames[0] = "Kashif"
	userNames[1] = "Shahzad"
	fmt.Println(userNames)

	userNames = append(userNames, "John")
	fmt.Println(userNames)

	userRatings := make(floatMap, 3)
	userRatings["go"] = 4.7
	userRatings["angular"] = 4.6
	userRatings["react"] = 4.8

	userRatings.output()

	//fmt.Println(userRatings)

	for index, value := range userNames {
		//...
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range userRatings {
		//...
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

}
