package main

import "fmt"

func main() {
	i:=1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	for {
		fmt.Println("loop")
		break
	}
}
