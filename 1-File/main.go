package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// read0 use os.open() and ioutil.Read()
func read0(fileName string) []byte {
	file, err := os.Open("test.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := make([]byte, 10)
	buff := make([]byte)
	for {
		n, err := file.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		result = append(result, buff[:n]...)
	return result
}

// read1 use os.open and ioutil.ReadAll()
func read1(fileName string) []byte {
	file, err := os.Open("test.json")
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

func read2(fileName string) []byte {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return f
}

func main () {
	fileName := "test.json"
	buffer := read0(fileName)
	fmt.Printf("Use Read0 Func %v\n", string(buffer))

	buffer1 := read1(fileName)
	fmt.Printf("Use Read1 Func %v\n", string(buffer1))

	buffer2 := read2(fileName)
	fmt.Printf("Use Read2 Func %v\n", string(buffer2))
}
