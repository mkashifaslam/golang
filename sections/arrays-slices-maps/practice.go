package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.

	hobbies := [3]string{"watching movies", "travel world", "reading tech articles"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	firstHobby := hobbies[0]
	fmt.Println(firstHobby)

	//secondThirdHobbies := hobbies[1:3]
	//fmt.Println(secondThirdHobbies)

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)

	newHobbiesM1 := hobbies[0:2]
	fmt.Println(newHobbiesM1)

	newHobbiesM2 := hobbies[:2]
	fmt.Println(newHobbiesM2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.

	newHobbiesM2 = newHobbiesM2[1:3]
	fmt.Println(newHobbiesM2)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)

	courseGoals := []string{"basics", "apis"}
	fmt.Println(courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "cli"
	courseGoals = append(courseGoals, "desktop")
	fmt.Println(courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		{
			title: "computer",
			id:    "123",
			price: 43.28,
		}, {
			title: "laptop",
			id:    "456",
			price: 230.12,
		},
	}
	products = append(products, Product{
		title: "tablet",
		id:    "789",
		price: 64.54,
	})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
