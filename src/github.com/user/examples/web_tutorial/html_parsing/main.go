package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type htmlData struct {
	attribute string
	value     string
}

var hData htmlData

func main() {

	in, _ := os.Open("sample.html")
	hData = htmlData{}
	token := html.NewTokenizer(in)

	printAtoken(token)

	fmt.Println(hData)

}

func emitToken(data []byte) {

	hData.value = hData.value + string(data)
}

func printAtoken(token *html.Tokenizer) {
	depth := 0
	var attrVal []byte
	for {
		tt := token.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.TextToken:
			if depth > 0 {
				emitToken(token.Text())
			}

		case html.StartTagToken, html.EndTagToken:
			tn, _ := token.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken && depth < 1 {
					_, attrVal, _ = token.TagAttr()
					hData.attribute = string(attrVal)
					depth++
				} else {
					depth--
				}
			}
		}

	}
}
