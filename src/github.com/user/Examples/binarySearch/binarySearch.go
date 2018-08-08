package main

import (
	"fmt"
	"sort"
)

func main() {
	myarray := readArray()
	mysort(myarray)
	fmt.Print("Sorted array : ")
	fmt.Println(myarray)

	var data int
	fmt.Printf("Enter element to search : ")
	fmt.Scanf("%d", &data)
	bSearch(myarray, 0, len(myarray), data)
}

func readArray() []int {
	fmt.Printf("Enter n :")
	var n int
	fmt.Scanf("%d", &n)
	var myarray = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter %d element: ", i+1)
		fmt.Scanf("%d", &myarray[i])
	}
	fmt.Println(myarray)
	return myarray
}

func mysort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func bSearch(arr []int, initial, final, data int) {
	mid := (initial + final) / 2
	if final > initial {
		if arr[mid] == data {
			fmt.Printf("Element Found in %d", mid+1)
		} else if arr[mid] < data {
			bSearch(arr, mid, final, data)
		} else {
			bSearch(arr, initial, mid, data)
		}
	} else {
		fmt.Println("Element is not presented in array")
	}
}
