package main

import "fmt"

func main() {
	countvalue := []int{30, 12, 28}
	sum := 0
	for _, num := range countvalue {
		sum += num
	}
	fmt.Println("count summary of the values:", sum)
	for i, num := range countvalue {
		if num == 30 {
			fmt.Println("index value of 30:", i)
		}
	}
	fruits := map[string]string{"a": "apple", "o": "orange", "w": "watermelon", "b": "banana"}
	for k, v := range fruits {
		fmt.Printf("fruit names:%s-->%s\n", k, v)
	}
	for _, v := range fruits {
		fmt.Println("values:", v)
	}
	for k := range fruits {
		fmt.Println("keys:", k)
	}
}
