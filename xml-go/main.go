package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Social struct {
	Facebook string `xml:"facebook"`
	Twitter  string `xml:"twitter"`
}

type User struct {
	Name   string `xml:"name"`
	Job    string `xml:"job"`
	Age    int    `xml:"age"`
	Social Social `xml:"social"`
}

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

func main() {
	file, err := os.Open("users.xml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Opened users.xml")
	defer file.Close()

	value, _ := ioutil.ReadAll(file)

	var users Users

	xml.Unmarshal(value, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Printf("Hello, %s \n", users.Users[i].Name)
		fmt.Printf("You are %d years old \n", users.Users[i].Age)
		fmt.Printf("You are currently a %s \n", users.Users[i].Job)
	}
}
