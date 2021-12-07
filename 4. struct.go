package main

import(
	"fmt"
)

type user struct {
	name  string
	age int
}

func main() {
	//1. Declaration
	employee := user{"Duy", 23}
	employee2 := user{name: "Cuong", age: 23}
	fmt.Println("1:",employee)
	fmt.Println("1:",employee2)
	

	// 2. Struct Instantiation Using new Keyword
	// => variable is reference type
	employee3 := new(user)
	employee3.name = "Trang"
	fmt.Println("2:",employee3)


}