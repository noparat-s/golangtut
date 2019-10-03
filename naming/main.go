package main

import "fmt"

var (
	anInteger int
	aString   = "string"
	arrayInt  = []int{1, 2, 3}
	arrayInt2 = [3]int{4, 5, 6}
	arrayInt3 = [...]int{7, 8, 9}
	intnull   []int
	mapInt    map[string]int
)

/*
pointer
slice
map
channel
function
interface
*/
func main() {
	fmt.Println(anInteger)
	fmt.Println(aString)
	fmt.Println(arrayInt)
	fmt.Println(arrayInt2)
	fmt.Println(arrayInt3)
	fmt.Println(intnull)
	fmt.Println(mapInt)

	fmt.Printf("%v\n", anInteger)
	fmt.Printf("%v\n", aString)
	fmt.Printf("%v\n", arrayInt)
	fmt.Printf("%v\n", arrayInt2)
	fmt.Printf("%v\n", arrayInt3)
	fmt.Printf("%v\n", intnull)
	fmt.Printf("%v\n", mapInt)

	if intnull == nil {
		fmt.Println("nil Array")
	}
	if mapInt == nil {
		fmt.Println("nil map")
	}
}
