package main

// import "fmt"
var a int
func changeA(b int) {
	a = b
}
func multAB(a, b int) int {
	return a * b
}

// func doPanic(a int) {
// 	if a > 1 {
// 		doPanic(a - 1)
// 	}
// 	panic(a)
// }
func main() {
	for i := 0; i < 10; i++ {
		changeA(i)
		multAB(i, 2 * i)
	}
	// fmt.Println("I am debugging now")
	// doPanic(5)
}