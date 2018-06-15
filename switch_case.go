package main
import "fmt"

func main() {
	s := 2
	switch  {
	case s > 5:
		fmt.Println("00")
	case s==4 :
		fmt.Println("aa")
	case s==3:
		fmt.Println("bb")
		fallthrough
	case s==2,s==1:
		fmt.Println("cc")
	default:
		fmt.Println("dd")
	}
}

