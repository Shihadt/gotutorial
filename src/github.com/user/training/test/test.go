package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()

	c := make(chan int)

	for i := 0; i < 5; i++ {
		fmt.Println("Asking some questions..")
		go func(no int) {
			var j int
			fmt.Println("question : ", no)
			go fmt.Scanf("%d", &j)
			c <- j
		}(i)

		go func() {
			select {
			case <-timer.C:
				fmt.Println("time out")
			case <-c:
				fmt.Println("input taken : ", <-c)
			}
		}()
	}
}

func test() {

}
