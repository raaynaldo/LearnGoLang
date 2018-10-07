package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["Ray"] = "ray@gmail.com"
	// emails["Felicia"] = "felicia@gmail.com"
	// emails["Brian"] = "brian@gmail.com"

	// Declare map and add kv
	emails := map[string]string{"Ray": "ray@gmail.com", "Felicia": "felicia@gmail.com", "Brian": "brian@gmail.com"}
	emails["Tim"] = "tim@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Felicia"])

	// Delete from map
	delete(emails, "Ray")
	fmt.Println(emails)

}
