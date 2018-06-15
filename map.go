package main
import "fmt"
import "reflect"

func createMapLiteral() {
	// syntax of map type
	var mm map[string]int
	fmt.Println(reflect.TypeOf(mm)) // map[string]int
	
	var mm2 = map[string]int{"a": 1, "b": 2}
	fmt.Println(mm2) // map[a:1 b:2]
}

func createMapMake() {
	var mp = make(map[string]int)
	mp["a"] = 1
	mp["b"] = 2
	fmt.Println(mp) // map[a:1 b:2]
}

func setAddGetDelete() {
	var mm = map[string]int{"a": 1, "b": 2}
	// set valude to existing key
	mm["a"] = 9
	fmt.Println(mm) // map[a:9 b:2]
	// if key don't exist, add it	
	mm["c"] = 8
	fmt.Println(mm) // map[a:9 b:2 c:8]
	
	// return 2 values. second is true if exist, else false
	var y, z = mm["c"]
	fmt.Println(y, z) // 8 true	
	var h1, h2 = mm["d"]
	fmt.Println(h1, h2)
}

func deleteKey() {
	var mm = map[string]int{"a": 1, "b": 2}
	// delete a key
	delete(mm, "a")
	fmt.Println(mm) // map[b:2]

	// if key doesn't exist, delete does nothing
	delete(mm, "9")
	fmt.Println(mm) // map[b:2]
}

func main () {
	createMapLiteral(); fmt.Println()
	createMapMake(); fmt.Println()
	setAddGetDelete(); fmt.Println()
	deleteKey(); fmt.Println()
}
