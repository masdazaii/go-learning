package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (p Person) GetName()  string {
	return p.Name
}

func main(){
	var eko Person
	eko.Name = "Anggit"

	SayHello(eko);
}