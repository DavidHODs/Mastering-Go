package chapterfour

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Parse() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide a filename")
		return
	}

	filename := arguments[1]

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var parsedData map[string]interface{}
	json.Unmarshal([]byte(fileData), &parsedData)

	for key, value := range parsedData {
		fmt.Println("key:", key, "value:", value)
	}
}