package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetJSONFile(filename string) []byte {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	jsonFile, err := os.Open(filename)
	defer jsonFile.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("File Successfully Opened")
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	return byteValue
}
