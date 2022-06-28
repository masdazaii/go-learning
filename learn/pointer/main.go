package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	//Pointer
	var address1 Address = Address{"Purwokerto", "Central Java", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Yogyakarta"

	*address2 = Address{"Malang", "East Java", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	fmt.Println(address4)

	//Pointer as function
	var alamat = Address{"Semarang", "Central Java", "Indonesia"}
	ChangeCountryToJepang(&alamat);
	fmt.Println(alamat)
}

func ChangeCountryToJepang( address *Address){
	address.Country = "Jepang"
}