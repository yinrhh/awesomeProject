package main

import "fmt"

func main() {

	var a [5]int
	a[4] = 100
	fmt.Println(a[4], len(a), a[3])

}
