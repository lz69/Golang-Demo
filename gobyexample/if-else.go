package main

import "fmt"

func main() {
	tip := "7 is even"
	if 7%2 != 0 { 
		tip = "7 is odd"
	}
	fmt.Println(tip)
}
