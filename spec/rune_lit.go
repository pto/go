package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("utf8.ValidRune('\\u00FF') is", utf8.ValidRune('\u00FF'))
	fmt.Println("utf8.ValidRune('\\xFF') is", utf8.ValidRune('\xFF'))
	fmt.Println("utf8.ValidString(\"\\u00FF\") is", utf8.ValidString("\u00FF"))
	fmt.Println("utf8.ValidString(\"\\xFF\") is", utf8.ValidString("\xFF"))
}
