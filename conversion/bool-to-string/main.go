package main

import (
	"fmt"
	"strconv"
)

func main()  {
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean ke string:", str)

	val, _ := strconv.ParseBool("true")
	fmt.Println("String ke boolean", val)
}
