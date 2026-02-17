package main

import "fmt"

func main()  {

	// defer fmt.Println("A: dieksekusi paling akhir")
	// fmt.Println("B: dieksekusi duluan")

	x := 1
	defer fmt.Println("defer-1, x =", x)
	x = 2
	defer fmt.Println("defer-2, x =", x)
	fmt.Println("body, x =", x)
}