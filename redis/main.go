package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	result, err := conn.Do("HMSET", "company:1", "name", "Amazon")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	company, err := redis.String(conn.Do("HGET", "company:1", "name"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The name of the company is %s\n", company)
}
