package main

import "fmt"

type PersegiPanjang struct {
	Panjang float64
	Lebar float64
}

// Method: menghitung luas
func (p PersegiPanjang) Luas() float64 {
	return p.Panjang * p.Lebar
}

func main()  {
	pp := PersegiPanjang{Panjang: 10, Lebar: 5}
	fmt.Println("Panjang:", pp.Panjang)
	fmt.Println("Lebar:", pp.Lebar)
	fmt.Println("Luas:", pp.Luas())	
}
