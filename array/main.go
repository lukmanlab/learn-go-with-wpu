package main

import "fmt"

func main()  {
	var angka [3]int = [3]int{10,20,30}
	fmt.Println(angka)
	fmt.Println(angka[1])

	angka[1] = 50
	fmt.Println("Ini adalah value index ke-1:", angka[1])

	// common function
	number := [5]int{10,20,30,40,50}

	fmt.Println("Jumlah element :", len(number)) // len(arr)
	
	for index, value := range number {
		fmt.Printf("Index ke-%v adalah: %v\n", index, value)
	}

	arr1 := [3]int{1,2,3}
	arr2 := [3]int{4,5,6}

	fmt.Println("arr1 == arr2", arr1 == arr2)
	fmt.Println("arr1 != arr2", arr1 != arr2)
}
