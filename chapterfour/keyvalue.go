package chapterfour

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement)

func add(k string, n myElement) bool {
	if k == "" {
		return false
	}

	if lookup(k) == nil {
		DATA[k] = n
		return true
	}

	return false
}

func deleteData(k string) bool {
	if lookup(k) != nil {
		delete(DATA, k)
		return true
	}

	return false
}

func lookup(k string) *myElement {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	} else {
		return nil
	}
}

func change(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func print() {
	for k, d := range DATA {
		fmt.Printf("key: %s value: %v/n", k, d )
	}
}

func Switch() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = s.TrimSpace(text)
		tokens := s.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			print()
		case "STOP":
			return
		case "DELETE":
			if !deleteData(tokens[1]) {
				fmt.Println("delete operation failed")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !add(tokens[1], n) {
				fmt.Println("add operation failed")
			}
		case "LOOKUP":
			n := lookup(tokens[1])
			if n != nil {
				fmt.Printf("%v/n", *n)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !change(tokens[1], n) {
				fmt.Println("update operation failed")
			}
		default:
			fmt.Println("unknown command - please try again")
		}
	}
}