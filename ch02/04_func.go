package main

import "fmt"

func main() {
	funcs := []string {
		"make",
		"len",
		"cap",
		"new",
		"append",
		"copy",
		"close",
		"delete",
		"complex",
		"real",
		"imag",
		"panic",
		"recover",
	}
	for _, val := range funcs {
		fmt.Println(val)
	}
}
