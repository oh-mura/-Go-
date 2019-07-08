package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const DISTANCE = 62100000

	// タイトルを先に出す。こうするしか思いつかない
	fmt.Printf("%-16v %-4v %-13v %v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=========================================")

	for count := 0; count < 10; count++ {
		// Select Spaceline
		var spaceLine string
		switch num := rand.Intn(3); num {
		case 0:
			spaceLine = "Space Adventures"
		case 1:
			spaceLine = "SpaceX"
		case 2:
			spaceLine = "Virgin Galactic"
		}

		// Select Speed & Price
		base := int16(rand.Intn(15)) // rand.Intnがintで結果返す
		speed := base + 16           // 16〜30の範囲
		price := base + 36           // 36〜50の範囲

		// 日数
		days := DISTANCE / (int64(speed) * 86400)

		// Select Trip Type
		var tripType string
		switch num := rand.Intn(2); num {
		case 0:
			tripType = "One-way"
		case 1:
			tripType = "Round-trip"
			price *= 2
		}

		fmt.Printf("%-16v %-4v %-12v %2v %-2v\n", spaceLine, days, tripType, "$", price)
	}

}
