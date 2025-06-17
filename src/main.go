package main

import (
	"fmt"
	"time"
)

func GetCurrentTime() {
	currentTime := time.Now()
	fmt.Println("Current time is =", currentTime)
}
