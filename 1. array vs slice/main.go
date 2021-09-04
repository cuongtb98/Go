package main

import (
	"fmt"
	"reflect"
)

func array() {
	// declare
	var array1 [3]int
	array2 := [...]int{4,5,6}
	array3 := [3]int{1,2,3}
	array4 := [3]*int{}
	array5 := [3]int{}

	// show
	fmt.Println("____________ BEGIN array ____________")
	fmt.Println("array1: ", array1, reflect.ValueOf(array3).Kind())
	array1[0] = 2
	fmt.Println("array1: ", array1)
	fmt.Println("array2: ", array2, reflect.ValueOf(array2).Kind())
	fmt.Println("array3: ", array3, &array3[0] , reflect.ValueOf(array3).Kind())
	fmt.Println("array4: ", array4)
	fmt.Println("array5: ", array5)
	fmt.Println("____________ END array   ____________")
}

func slice() {
	// datatype:  reference types
	// declare
	var slice1 []int
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 2, 3}
	slice4 := slice3
	
	// append 
	slice2 = append(slice2, 4)
	slice2 = append(slice2, []int{5,6}...)

	// copy
	slice6 := make([]int,2)
	slice7 := []int{4,5,6,7,8,9}
	

	// show
	fmt.Println("____________ BEGIN slice ____________")
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)
	slice4[0] = 2
	fmt.Println("slice3", slice3)
	fmt.Println("slice6", copy(slice6, slice3), slice6)
	fmt.Println("slice7", copy(slice7, slice3), slice7)
	fmt.Println("____________ END slice ____________")
}

func main() {
	array()
	slice()
}