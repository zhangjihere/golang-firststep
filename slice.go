package main
import "fmt"
import "reflect"

func literalExpression(){
	//array
	var aa = [3]int{9, 2,  6}
	// slice
	var ss = []int{9, 2, 6}

	fmt.Println(reflect.TypeOf(aa))	
	fmt.Println(reflect.TypeOf(ss))
}

func creatSliceMake() {
	var s = make([]int, 3)	// 3 items of int
	fmt.Println(s)		// [0 0 0]
	var s2 = make([]int, 3, 9)	// 3 item of int, capacity of 9
	s2[2] = 2
	fmt.Println(s2, len(s2))			// [0 0 0]
}

func takeSliceOfSlice() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(s) // [1 2 3 4 5 6 7]
	// take aslice of values. include first index, exclude second index
	fmt.Println(s[1:3]) // [2 3]
}

func append2Slice() {
	var ss = []int{3, 4, 5}
	var s2 = append(ss, 8, 9)
	fmt.Println(s2) // [3 4 5 6 8 9]
}

func nestedSlice(){
	// y is slice of slice of int
	var y = [][]int{{3, 4}, {7, 8, 9},{1, 2}}
	fmt.Println(y) // [[3 4] [7 8 9] [1 2]]

	var y2 = make([][]string, 2)
	y2[0] = []string{"a","b"}
	y2[1] = []string{"x", "y", "z"}
	fmt.Println(y2) // [[a b] [x y z]]
}

func zeroValueOfSliceIsNil() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")	
	}
}

func loopThruSlice() {
	// slice
	var s = []int {9, 2, 8, 61}
	for i:=0; i< len(s); i++ {
		fmt.Println(i, s[i])
	}
}

func main() {
	literalExpression(); fmt.Println("")
	creatSliceMake(); fmt.Println("")
	takeSliceOfSlice(); fmt.Println("")
	append2Slice(); fmt.Println("")
	nestedSlice(); fmt.Println("")
	zeroValueOfSliceIsNil(); fmt.Println("")
	loopThruSlice()
}
