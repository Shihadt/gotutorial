package main

import (
	"fmt"
)

type mytype []string

func (m mytype) myfunc() {
	for i, j := range m {
		fmt.Println(i, j)
	}
}
