package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter number : ")
	fmt.Scanf("%d", &n)

	if isAmstrong(n) {
		fmt.Print("Number is amstrong")
	} else {
		fmt.Print("Number is not amstrong.")
	}
}
