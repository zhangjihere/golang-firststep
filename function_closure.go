package main
import "fmt"

// return a closure
func make_add_1() func(int) int {
	var i = 0;
	
	var gg = func(x int) int {
		i +=x
		return i
	}
	
	// the returned function gg, is called a "closure"
	return gg
}

func main() {
	var f = make_add_1()
	
	fmt.Println(f(1)) // 1
	fmt.Println(f(2)) // 2
	fmt.Println(f(3)) // 3
}
