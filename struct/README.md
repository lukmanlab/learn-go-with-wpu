# Struct
- Struct (structure) adalah tipe data di Go yang memungkinkan kita mengelompokkan beberapa data dengan tipe berbeda ke dalam satu kesatuan.
- Bisa dianggap seperti "kerangka" atau "blueprint" untuk membuat objek dengan properti tertentu
- Mirip dengan class di bahasa lain, hanya saja Go tidak memiliki class dan inheritance, melainkan struct + interface
- Struct merupakan pondasi untuk membangun objek di Go--sederhana, tapi bisa berkembang jadi sistem yang kompleks.

# Fungsi
1. Mengelompokkan data yang saling berkaitan.
   - Misalnya, data tentang User bisa terdiri dari Name, Email, Age
2. Membuat representasi objek dunia nyata.
   - Misalnya: Mobil, Mahasiswa, Buku, dll.
3. Mempermudah modularisasi kode agar lebih rapi, terstruktur, dan mudah dibaca.

## Sample
```golang
package main

type User struct {
  Name string
  Email string
  Age int
}
```

# Catatan Penting
- Struct tidak memiliki inheritance, tetapi bisa menggunakan composition (menyisipkan struct lain didalam struct)
- Struct sering digunakan bersama dengan method agar bisa melakukan aksi (mirip oop tapi versi Go)


## Sample - composition
```golang
type Address struct {
  City, Province string
}

// Struct User menggunakan composition
type User struct {
  Name string
  Age int
  Address // embedded struct (composition)
}
```

# Struct Method
Method adalah fungsi yang menempel pada struct. Dengan method, kita bisa memberikan perilaku (behavior) ke sebuah struct, mirip seperti OOP.
