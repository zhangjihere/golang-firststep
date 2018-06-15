package main
import "fmt"

// ff is my function name
func ff(x int, y int) int {
	return x + y
}

func ff2(x,y int) int {
	return x*y
}

func main() {
	fmt.Println(ff(3,4))
	fmt.Println(ff2(3,4))
}
// print 7
// print 12
