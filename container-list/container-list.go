package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Rifan")
	data.PushBack("Nur")
	data.PushBack("Muhammad")

	// jika ingin ambil data setelah paling depan
	// data.Front().Next()

	// jika ingin ambil data sebelumnya dari paling depan
	// data.Front().Prev()

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	// Dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// Dari belakang ke depan
	// for element := data.Back(); element != nil; element = element.Prev() {
	// 	fmt.Println(element.Value)
	// }
}
