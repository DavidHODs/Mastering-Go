package chapterfive

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randoms(min, max int) int {
	return rand.Intn(max-min) + min
}

func Randoms() {
	MIN := 0
	MAX := 94
	var LENGTH int64 = 8
	arguments := os.Args

	switch len(arguments) {
	case 2:
		LENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
	default:
		fmt.Println("Using default values")
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	startChar := "!"
	var i int64 = 1
	
	for {
		myRand := randoms(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)

		if i == LENGTH {
			break
		}
		i++
	}

	fmt.Println()
}