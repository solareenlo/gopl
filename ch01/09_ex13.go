package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	num := make([]string, 10000, 10000)
	strings.Join(num[:], " ")
	fmt.Println(time.Since(start).Seconds())
}
