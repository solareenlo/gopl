package main

import "fmt"

func main() {
	consts := []string{
		"true",
		"false",
		"iota",
		"nil",
	}
	for _, val := range consts {
		fmt.Println(val)
	}
}
