package chaptertwo

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
)

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Println(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}

	fmt.Println()
}


func d3() {
	for i := 3; i < 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

// Defer shows the defer capability in Go
func Defer() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}

var LOGFILE = "tmp/mGo.log"

// DeferLog shows defer in loggin situations
func DeferLog(aLog *log.Logger) {
	aLog.Println("----FUCTION TWO-----")
	defer aLog.Println("Function two-------")

	for i := 10; i > 0; i-- {
		aLog.Println(i)
	} 
}

func Version() {
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]

	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go version 1.8 or higher")
		return
	}

	fmt.Println("You are using Go version 1.8 or higher")
	fmt.Println(myVersion)
}