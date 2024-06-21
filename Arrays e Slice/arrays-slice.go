package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [10]string
	array1[0] = "Flamengo"
	fmt.Println(array1)

	array2 := [5]string{"Flamengo", "Real madrid", "Bufallo bills", "Oilers", "Lakers"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array3)

	slice := []int{23, 56, 89, 45, 0, 23, 45, 50, 56}
	fmt.Println(slice)

	//Reflete, literalmente o tipo das variaveis
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 19)
	fmt.Println(slice)

	slice2 := array2[1:4]
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(slice2)
}
