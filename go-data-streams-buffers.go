package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//create a buffer to hold JSON data
	buf := new(bytes.Buffer)
	//create JSON encoder for `buf`
	bufEncoder := json.NewEncoder(buf)

	//encode JSON from `Person` struct
	bufEncoder.Encode(Person{Name: "Arpit", Age: 24})
	bufEncoder.Encode(Person{Name: "Mohit", Age: 22})

	//print contents of `buf`
	fmt.Println(buf)

	//create JSON decoder using `buf`
	decoder := json.NewDecoder(buf)

	//create `Person` struct to hold decoded data
	var arpit, mohit Person

	//decode JSON from `decoder` one line at a time
	decoder.Decode(&arpit)
	decoder.Decode(&mohit)

	//see the values of `arpit` and `test`
	fmt.Printf("Arpit: %#v\n", arpit)
	fmt.Printf("test: %#v\n", mohit)
}
