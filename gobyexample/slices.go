package main

import "fmt"

func main() {
	
	s:=make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s, s[2], len(s), append(s, "e", "f" ), s[1:3])
}
