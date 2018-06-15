package main
import "fmt"

//return a function
func ff(x int) func(int) int {
	var gg = func(y int) int { return y+x}
	return gg
}

func main() {
	// assign a function to a variable
	var hh = func() int {
		return 3
	}
	// nested function
	var hh2 = func(x int) int {
		return x + 1	
	}
	fmt.Println(hh()) // 3
	fmt.Println(hh2(3)) // 4		

	// return a function 
	var ret = ff(2)(3)
	fmt.Println(ret)
}
