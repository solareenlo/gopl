package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := ""
	num := make([]string, 100000, 100000)
	for _, i := range num[:] {
		s += i
		fmt.Println(s)
	}
	fmt.Println(time.Since(start).Seconds())
}
