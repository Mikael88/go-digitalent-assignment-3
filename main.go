package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Status struct {
	Water int
	Wind int
}

func randomInt(min, max int) int {
	return rand.Intn(max+1) + min
}

func generateRandomData() Status {
	return Status{
		Water: randomInt(1, 100),
		Wind: randomInt(1, 100),
	}
}

func main() {
	ticker := time.NewTicker(15 * time.Second)
	for range ticker.C {
		data := generateRandomData()
		
		fmt.Println("Water: ", data.Water)
		fmt.Println("Wind: ", data.Wind)

		waterStatus := "Aman"
		if data.Water > 8 {
			waterStatus = "Bahaya"
		} else if data.Water >= 6 {
			waterStatus = "Siaga"
		}

		windStatus := "Aman"
		if data.Wind > 15 {
			windStatus = "Bahaya"
		} else if data.Wind >= 7 {
			windStatus = "Siaga"
		}

		fmt.Println("Water status: ", waterStatus)
		fmt.Println("Wind status: ", windStatus)
		fmt.Println()
	}
}
