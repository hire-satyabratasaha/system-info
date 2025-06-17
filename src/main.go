package src

import (
	"fmt"
	"time"
)

func GetCurrentTime() (string, int) {
	timex := time.Now().String()
	_, _, day := time.Now().Date()
	return timex, day
}

func GetCurrentDate() {
	year, month, day := time.Now().Date()
	fmt.Println("Current time is =", year, month, day)
}
