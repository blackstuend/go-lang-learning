package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Friend struct {
	Name string
	Age  int
}

type User struct {
	ID        int `json:"id"`
	FirstName string
	LastName  string
	Age       int
	Friends   []Friend
}

func openJsonFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	buff, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return buff
}

const (
	fileName string = "Friend.json"
)

func main() {
	staticJson()
	fileJson()
	receiveAllData()

}

func staticJson() {
	var jsonString string = `{
		"id":0,
		"FirstName":"Lin",
		"LastName":"Vincent",
		"age":10
	}`
	var user User
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%+v\n", user)

	var u map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &u)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%+v\n", u)

}

func fileJson() {
	jsonBuffer := openJsonFile(fileName)

	var user User
	err := json.Unmarshal([]byte(jsonBuffer), &user)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%+v\n", user)
}

func receiveAllData() {
	var user map[string]interface{}
	jsonBuffer := openJsonFile(fileName)
	err := json.Unmarshal([]byte(jsonBuffer), &user)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%+v\n", user)
}
