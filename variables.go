package main
import "fmt"

// package level variable
var c, d, e int

var i2, j = 3, "no"

func main() {
	// function level variable
	var i int
 	fmt.Println(i, c, d, e) // 0 0 0 0
	fmt.Println(i2, j) // 3 no
	
	k := 4
	fmt.Println(k)

	// initial value, zero values
	var n int
	var m float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", n, m, b, s)
	// 0 0 false ""
}

