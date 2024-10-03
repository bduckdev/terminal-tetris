package main

import (
	"fmt"
	"time"
)

func main() {
	daysSpentDucking := 0

	fmt.Println("Days spent ducking: ", daysSpentDucking)

	for {
		time.Sleep(24 * time.Hour)
		daysSpentDucking++
		fmt.Println("Days spent ducking: ", daysSpentDucking)
	}
}
