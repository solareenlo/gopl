package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	s := ""
	for i := 0; i < 10000; i++ {
		s += strconv.Itoa(i)
	}
	fmt.Println(time.Since(start).Seconds())
}
