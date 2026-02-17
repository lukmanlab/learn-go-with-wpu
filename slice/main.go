package main

import "fmt"

func main()  {
	arr := [6]int{10,20,30,40,50,60}

	s1 := arr[:] // ambil seluruh element
	fmt.Println("Ini adalah slice 1 :", s1)

	s2 := arr[:3]
	fmt.Println("Ini adalah slice 2 :", s2)

	s3 := arr[2:]
	fmt.Println("Ini adalah slice 3:", s3)

	s4 := arr[1:4]
	fmt.Println("Ini adalah slice 4:", s4)

	// next

	s := make([]int, 3, 5)
	fmt.Println(s)
	fmt.Println("Length :", len(s))
	fmt.Println("Capacity :", cap(s))

	s = append(s, 10, 20)
	fmt.Println("Setelah di append :", s)

	// next
	slice2 := make([]int, 3)
	slice3 := copy(slice2, s)
	
	fmt.Println("Copied :", slice2)
	fmt.Println("jumlah elemen yang tersalin :", slice3)

	angka := []int{1,2,3,4,5}
	slice4 := angka[1:4]
	fmt.Println("Slice 4:", slice4)
}
