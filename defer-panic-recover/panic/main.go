package main

import "fmt"

func main()  {
	defer fmt.Println("defer ini tetap jalan sebelum program mati")
	fmt.Println("Sebelum panic")
	panic("ada sesuatu yang fatal")
	fmt.Println("Setelah panic")
}
