package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	ticker := time.NewTicker(time.Second * 5)
	random := uuid.New()
	var currentTime time.Time

	for {
		currentTime = <-ticker.C
		fmt.Printf("%s: %s\n", currentTime.Format("2006-01-02T15:04:05.000Z"), random)
	}
}
