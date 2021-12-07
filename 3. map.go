package main

import(
	"fmt"
)

func myMap(){
	//1. Map initialization
	var employee = map[string]int{"Cuong": 23, "Duy": 22}
	fmt.Println("1:",employee)

	//2. Empty Map declaration
	var employee2 = map[string]int{}
	fmt.Println("2:",employee2)

	//3. Map declaration using make function
	var employee3 = make(map[string]int)
	employee4:= make(map[string]int)
	employee3["Cuong"] = 24
	employee4["Cuong"] = 25
	fmt.Println("3:",employee3)
	fmt.Println("3:",employee4)

	//4. Map Length
	fmt.Println("4:",len(employee))

	// 5. Accessing Items
	fmt.Println("5:",employee["Cuong"])

	//6. Adding Items
	employee["Trang"] = 23
	fmt.Println("6:",employee)

	// 7. Update Values
	employee["Trang"] = 22
	fmt.Println("7:",employee)

	// 8. Delete Items
	delete(employee, "Trang")
	fmt.Println("8:",employee)

	// 9. Sort Map Keys
	


}

func main() {
	myMap()
}