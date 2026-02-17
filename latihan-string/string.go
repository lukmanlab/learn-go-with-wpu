package main

import (
	"fmt"
	"strings"
)

func main() {
	nama := "Triady"
	pesan := "Selamat datang di aplikasi kami!"

	paragraf := `Halo, ini adalah contoh multi line
string di Go.`

	fmt.Println("Nama: ", nama)
	fmt.Println("Pesan: ", pesan)
	fmt.Println("Paragraf: ", paragraf)


	text := "Halo, Guys!"
	
	fmt.Println("Lowercase:", strings.ToLower(text))

	fmt.Println("Uppercase:", strings.ToUpper(text))

	fmt.Println("StartWith Halo?", strings.HasPrefix(text, "Halo"))

	fmt.Println("Contain Guys?", strings.Contains(text, "Guys"))
	
	parts := strings.Split("apel,banana,cerry", ",")
	fmt.Println("Split: ", parts)

	newText := strings.ReplaceAll(text, "Guys", "Golang")
	fmt.Println("Replace: ", newText)
}
