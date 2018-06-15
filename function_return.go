package main
import "fmt"

// return2 2 values
func ff(x, y int) (int, int) {
	return x, y
}

//return values can be named.
//If so, they are treated as variables defined at the top of the function
func gg(n int) (m int) {
	m = n + 1
	return m
}

func main() {
	var a, b = ff(3,4)
	fmt.Println(a, b)
	var c = gg(5)
	fmt.Println(c)
}
// prints 3 4
// prints 6
