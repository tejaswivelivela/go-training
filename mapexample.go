package main

import "fmt"

func main() {
	names := make(map[string]string)
	names["roll1"] = "teju"
	names["roll2"] = "sai"
	names["roll3"] = "chinna"
	v1, exists := names["roll1"]
	fmt.Println(" roll1 exists?", exists)
	fmt.Println("value of roll1:", v1)
	delete(names, "roll3")
	fmt.Println("student names:", names)
	fmt.Println("lenght of the map:", len(names))

	role3, exists := names["roll3"]
	fmt.Println("roll3:", role3)
	fmt.Println("exists?", exists)
	initials := map[string]string{"teju": "veliv", "sai": "kota"}
	fmt.Println("initials of names:", initials)
	fmt.Println("length of the map:", len(initials))

	//anyMap := make(map[string]interface)
	// anyMap["name"] = "Sai"
	//anyMap["ID"] = 1234
	//fmt.Println("anymap:",anyMap)
}
