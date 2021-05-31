package main

import "fmt"

type Guitarist interface {
	PlayGuitar()
}

func (g Guitarist) PlayGuitar() {
	fmt.Printf("Hello, Guitarist: %s", g.Name)
}

func main() {

}
