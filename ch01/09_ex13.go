package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	num := make([]string, 100000, 100000)
	fmt.Println(strings.Join(num[:], " "))
	fmt.Println(time.Since(start).Seconds())
}
