package main
import "fmt"

func main() {
	var a [2]string
	a[0] = "cat"
	a[1] = "dog"
	fmt.Println(a) 		// [cat dog]
	fmt.Println(len(a))	//2
	
	//array declare and init
	// literal expression to init array slots
	var x = [4]int{5, 3, 2, 9}
	fmt.Println(x) // [5 3 2 9]
}
