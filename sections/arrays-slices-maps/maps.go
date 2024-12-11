package main

import "fmt"

func Maps() {
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://www.aws.com",
	}
	fmt.Println(websites)

	websites["Linkedin"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
