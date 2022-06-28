package main

import (
	"fmt"
	"learn/constant"
	"learn/database"
	"learn/helper"
)

func main() {

	// Access Modifier
	helper.SayHello("anggit")
	fmt.Println(constant.API_URL)

	//Package Initialization
	result := database.TestConnection()
	fmt.Println(result)
}