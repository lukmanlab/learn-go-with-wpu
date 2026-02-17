package main

import "fmt"

func tanganiPanic()  {
	if r := recover(); r != nil {
		fmt.Println("Terjadi panic, tapi sudah ditangani: ", r)
	}
}

func bagi(a, b int)  {
	defer tanganiPanic()
	fmt.Printf("Membagi %d dengan %d\n", a, b)
	hasil := a / b
	fmt.Println("Hasil ", hasil)
}

func main()  {
	fmt.Println("Program mulai")
	bagi(10, 2)
	bagi(5, 0)
	fmt.Println("Program selesai dengan aman")
}