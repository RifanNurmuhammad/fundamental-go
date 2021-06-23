package main

import (
	"fmt"
)

// ContextKey type
type ContextKey string

func main() {
	var abc string
	newContext(&abc)
	fmt.Println(abc)
	fmt.Println("__________________________")
}

func newContext(abc *string) {
	var newabc string
	newabc = "abc masuk"
	abc = &newabc
}
