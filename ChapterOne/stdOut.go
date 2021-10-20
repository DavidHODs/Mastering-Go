package chapterone

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Writer demonstrates the io package
func Writer() {
	myString := ""
	arguments := os.Args
	myStringByte := []byte("An arguement in byte")
	
	if len(arguments) == 1 {
		myString = "Please give me one argument"
	} else {
		myString = arguments[1] + " " + arguments[2]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
	io.Writer.Write(os.Stdout, myStringByte)
}

// Reading from standard input
func Reader() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

// CommandLine demonstrates working with command line arguments
func CommandLine() {
	if len(os.Args) == 1 {
		fmt.Println("Please give me one or more floats")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[1], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}