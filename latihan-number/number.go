package main

import "fmt"

func main() {
	// signed integers
	var i8 int8 = 127
	var i16 int16 = 21767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = -100 // ukuran tergantung arsitektur (32-bit atau 64-bit)

	// unsigned integers
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 100 // tergantung arsitektur

	fmt.Println("Signet integers:")
	fmt.Printf("int6 : %v\n", i8)
	fmt.Printf("int6 : %v\n", i16)
	fmt.Printf("int6 : %v\n", i32)
	fmt.Printf("int6 : %v\n", i64)
	fmt.Printf("int6 : %v\n", i)

	fmt.Println("\nUnsigned Integers:")
	fmt.Printf("uint8 : %v\n", u8)
	fmt.Printf("uint8 : %v\n", u16)
	fmt.Printf("uint8 : %v\n", u32)
	fmt.Printf("uint8 : %v\n", u64)
	fmt.Printf("uint8 : %v\n", u)
}
