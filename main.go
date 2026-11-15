package main

import (
	"fmt"
	"go-repeat-1/internal/file"
	"go-repeat-1/internal/text"
)

func main() {
	content, err := file.ReadFile("internal/testdata/input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(text.ToUpperCase(content))
	fmt.Println(content)
	fmt.Println(text.CountWords(content))

	err = file.WriteFile("internal/testdata/output.txt", "Hello World")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
}
