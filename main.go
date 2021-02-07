package main

import (
	"fmt"

	qt "github.com/tinjaw/stringmod/quotes"
	str "github.com/tinjaw/stringmod/strings"
)

func main() {
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e) // 3 2

	fmt.Println(qt.GetEmoji("turtle"))
}
