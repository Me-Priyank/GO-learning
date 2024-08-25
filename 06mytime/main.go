package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time handeling part")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05")) // We always have to use the exact instance as a standard to get formatted data

	createDate := time.Date(2025, time.February, 10, 12, 00, 00, 00, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 Monday 15:04:05"))
}
