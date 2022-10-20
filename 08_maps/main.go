package main

import "fmt"

func main() {
	// Define map
	email := make(map[string]string)
	email["Gaurav"] = "gauravshinde@gmail.com"

	// Decalre map and add kv
	emails := map[string]string{"Gaurav": "gauravshinde@gmail.com", "saurabh": "saurabh@gmail.com"}
	emails["Gaurav"] = "gauravshinde@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["saurabh"])

	// Delete from map
	delete(emails, "saurabh")
	fmt.Println(emails)
}
