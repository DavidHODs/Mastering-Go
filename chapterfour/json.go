package chapterfour

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	err = json.NewDecoder(in).Decode(key)
	if err != nil {
		return err
	}

	return nil
}

func ReadJson() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide a filename")
		return
	}

	filename := arguments[1]

	var myRecord Record

	err := loadJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}

func encodeJson(filename *os.File, key interface{}) {
	err := json.NewEncoder(filename).Encode(key)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func SaveJSON() {
	myRecord := &Record{
		Name: "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{{Mobile: true, Number: "1234-567"},
			{Mobile: true, Number: "1234-abcd"},
			{Mobile: false, Number: "abcc-567"},
			},
	}

	encodeJson(os.Stdout, myRecord)
}

func Marshal() {
	myRecord := Record{
		Name: "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{{Mobile: true, Number: "1234-567"},
			{Mobile: true, Number: "1234-abcd"},
			{Mobile: false, Number: "abcc-567"},
			},
	}

	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(rec))

	var unRec Record

	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(unRec)
}