package main

import (
	"fmt"
	"time"
)

func GetCurrentTime() {
	currentTime := time.Now()
	fmt.Println("Current time is =", currentTime)
}

func GetCurrentDate() {
	year, month, day := time.Now().Date()
	fmt.Println("Current time is =", year, month, day)
}
