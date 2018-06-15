package main
import "fmt"

func ff(x int) int {
	if x<0 {
		return -1
	} else if x > 0 {
		return 1
	} else {
		return 0
	}
}

func fd(x int) string {
	// if statement can start with a short statement, e.g. var declaration
	if i := -2; x < i {
		return "no"
	}
	return "yes"

}

// There's no "if expression", such as Javascripts (test ? expr1 : expr2)
func ternary(x int) string {
	//return (x ? fmt.Println(x) : fmt.Println(-x))
	return "empty"
}

func main() {
	fmt.Println(ff(-1), ff(2), ff(0))
	fmt.Println(fd(-3),fd(3))
	//fmt.Println(ternary(-3))
}

// -1 1 0
