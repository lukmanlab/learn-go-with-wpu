package main

import "fmt"

func cleanUp()  {
	fmt.Println("Cleanup: Menutup resource...")
}

func bacaConfig(namaFile string)  {
	defer cleanUp()

	defer func()  {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error:", r)
		}
	}()

	if namaFile == "" {
		panic("Nama file konfigurasi tidak boleh kosong")
	}

	fmt.Println("Membaca konfigurasi dari", namaFile)
}

func main()  {
	bacaConfig("")
	fmt.Println("Program tetap berjalan walaupun setelah panic")
}