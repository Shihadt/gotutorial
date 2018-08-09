package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	c := make(chan string)
	for i := 0; i < 5; i++ {
		go test(c)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i, " : channel: ", <-c)
	}
}

func test(c chan string) {
	var i int
	fmt.Scanf("%d", &i)
	c <- strconv.Itoa(i)
}
func myMain() {
	filename := flag.String("file", "problems.csv", "set File to read")
	//duration := flag.Duration("time", 30*time.Second, "Sets time duration")
	flag.Parse()

	r := readCsv(*filename)
	correct, wrong := printCsv(r)

	fmt.Println("Correct: ", correct, "\nWrong: ", wrong)
}

func readCsv(filename string) *csv.Reader {
	in, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	r := csv.NewReader(in)
	return r
}
func printCsv(r *csv.Reader) (int, int) {
	correct, wrong := 0, 0

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			println(err.Error())
		} else {
			correct, wrong = askAndcheckQuestion(record)
		}
	}
	return correct, wrong
}

func askAndcheckQuestion(record []string) (correct, wrong int) {
	correct, wrong = 0, 0
	var input int64
	input = 0

	fmt.Println("What is ", record[0], "?")
	fmt.Scanf("%d", &input)

	answer, err := strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}

	if input == answer {
		correct++
	} else {
		wrong++
	}
	return correct, wrong
}
