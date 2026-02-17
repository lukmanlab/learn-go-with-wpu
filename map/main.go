package main

import "fmt"

func main()  {
	mahasiswa := map[string]string{
		"nama": "ahmad",
		"kelas": "A",
	}
	// fmt.Println(mahasiswa)

	fmt.Println("Nama :", mahasiswa["nama"])
	fmt.Println("Kelas :", mahasiswa["kelas"])

	mahasiswa["jurusan"] = "informatika"
	fmt.Println("Jurusan", mahasiswa["jurusan"])
}


