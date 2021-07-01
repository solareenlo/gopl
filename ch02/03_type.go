package main

import "fmt"

func main() {
	types := []string {
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64",
		"complex128",
		"complex64",
		"bool",
		"byte",
		"rune",
		"string",
		"error",
	}
	for _, val := range types {
		fmt.Println(val)
	}
}
