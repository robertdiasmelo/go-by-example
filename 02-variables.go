package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var g string
	fmt.Println(g)

	var h, i = "banana", 50
	fmt.Println(h, i)

	// ERROR - Both j and k variables are 'int values' by inference
	// var j , k int = "banana", 10

}
