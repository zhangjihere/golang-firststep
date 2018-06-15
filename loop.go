package main
import "fmt"

// example of for loop
func main() {
	// scope of i is within the curly bracket{}
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Printf("%s, %q\n","=======","======")

	var ss = 1
	for ss <=3 {
		fmt.Println(ss)
		ss++
	}
	fmt.Printf("%s, %q\n","=======","======")
	
	uu := 1
	// infinite loop, use brak to exit
	for {
		uu++
		fmt.Println(uu)
		if uu >= 4 {
			break;
		}
	}
	fmt.Printf("%s, %q\n","=======","======")
	
	//break&continue loop
	for i := 1; i<=5; i++ {
		if i==4 {
			fmt.Println("  ==> i==4 break")
			break
		}
		if i==2 {
			fmt.Println("  ==> i==2 coninue")
			continue
		}
		fmt.Println(i)
	}

}





