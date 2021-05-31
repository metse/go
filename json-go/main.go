package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type User struct {
	Name   string `json:"name"`
	Job    string `json:"job"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Users struct {
	Users []User `json:"users"`
}

func main() {
	file, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Opened users.json file")
	defer file.Close()

	value, _ := ioutil.ReadAll(file)

	var users Users

	json.Unmarshal(value, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Job: " + users.Users[i].Job)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
