package main

import "fmt"
import "reflect"

func findValueType() {
	var nn = 3
 	var ff = 3.123
	// find the type of a value
	fmt.Println(reflect.TypeOf(nn)) //int
	fmt.Println(reflect.TypeOf(ff)) //float64	
}

func typeConversions() {
	var i int = 42
	var f = float64(i)
	fmt.Println(reflect.TypeOf(f)) // float64
}

func typeInference() {
	var i = 42			// int
	var f = 3.141592653589793	// float64
	var g = 1 + 0.5i		// complex128
	
	fmt.Println(reflect.TypeOf(i))	//int
	fmt.Println(reflect.TypeOf(f))	//float64
	fmt.Println(reflect.TypeOf(g))	//complex128
}

func main() {
	findValueType(); fmt.Println("")
	typeConversions(); fmt.Println("")
	typeInference() 
}	
