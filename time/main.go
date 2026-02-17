package main

import (
	"fmt"
	"time"
)

func convertToTimeZone(t time.Time, zone string) time.Time {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		fmt.Println("Error memuat lokasi:", err)
		return t
	}
	return t.In(loc)
} 

func main()  {
	nowUTC := time.Now().UTC()
	fmt.Println("Waktu UTC", nowUTC)

	jakartaTime := convertToTimeZone(nowUTC, "Asia/Jakarta")
	fmt.Println("Waktu Jakarta:", jakartaTime)

	newYorkTime := convertToTimeZone(nowUTC, "America/New_York")
	fmt.Println("Waktu New York:", newYorkTime)

	fmt.Println("Disimpan ke databasa (UTC)", nowUTC)
}