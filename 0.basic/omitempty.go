package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type PP struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {

	file, err := ioutil.ReadFile("./person.json")

	if err != nil {
		fmt.Printf("readfile fail err:%v", err)
	}

	p := PP{}

	json.Unmarshal(file, &p)

	fmt.Printf("person is %+v\n", p)
}
