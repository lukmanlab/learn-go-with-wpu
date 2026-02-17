package main

import "fmt"

func ubahNama(nama *string)  { // flow 3: pointer nama diisi oleh alamat jl. merdeka
	*nama = "Budi" // ubah langsung ke alamat memory aslinya

	// flow 4: jl. merdeka diganti nilai / value nya menjadi = Budi
}



func main()  {
	nama := "Andi" // flow 1: misalkan variable ini beralamat di jl. merdeka
	ubahNama(&nama) // flow 2: berikan alamat jl. merdeka ke fungsi ubahNama

	fmt.Println(nama) // output: Budi
}
