package main

import (
	"fmt"
)

type Address struct {
	City, Country, Province string
}

func main() {

}

// Man struct
type Man struct {
	Name string
}

// Married func changes if married
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func changeCountryToJapan(address *Address) {
	address.Country = "Japan"
}

func example1() {
	// CONTOH 1 tidak merubah reference data
	address1 := Address{"Bogor", "Indonesia", "Jawa Barat"}
	address2 := address1

	address2.Country = "Belanda"

	fmt.Println(address1)
	fmt.Println(address2)
}

func example2() {

	// CONTOH 2 ubah reference address data
	var address3 Address = Address{"Bogor", "Indonesia", "Jawa Barat"}
	var address4 *Address = &address3
	address4.Country = "Belanda"

	fmt.Println(address3)
	fmt.Println(address4)
}

func example3() {

	// CONTOH 3 tidak berubah tanpa operator (*)
	var address5 Address = Address{"Bogor", "Indonesia", "Jawa Barat"}
	var address6 *Address = &address5
	address6.Country = "Belanda"

	address6 = &Address{"Jakarta", "Indonesia", "DKI"}
	fmt.Println(address5)
	fmt.Println(address6)

}

func example4() {
	// CONTOH 4 tidak berubah reference dengan operator (*)
	var address7 Address = Address{"Bogor", "Indonesia", "Jawa Barat"}
	var address8 *Address = &address7
	var address9 *Address = &address7
	*address8 = Address{"Bandung", "Indonesia", "Jabar"}
	fmt.Println(address7)
	fmt.Println(address8)
	fmt.Println(address9)
}

func example5() {
	// CONTOH 5 create kosong
	var address10 *Address = new(Address)
	address10.Country = "Argentina"
	fmt.Println(address10)
}

func exmple6() {

	// CONTOH 6 change reference with call function
	var address11 Address = Address{"Bogor", "Indonesia", "Jawa Barat"}
	changeCountryToJapan(&address11)

	fmt.Println(address11)
}

func exmple7() {

	// CONTOH 7 pointer method menggunakan (*) di method / struct function
	budi := Man{"budi"}
	budi.Married()
	fmt.Println(budi)
}
