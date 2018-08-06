package main

import (
	"fmt"
	"strconv"
)

func main() {
	array := []string{"some", "and other", "also"}
	for i, arr := range array {
		fmt.Printf(strconv.Itoa(i) + arr)
	}
}
