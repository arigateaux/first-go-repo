package main

import (
	"fmt"

	"github.com/arigateaux/animals"
)

type Speaker interface {
	Speaks() string
}

func Perform(a Speaker) string {
	return a.Speaks()
}

func main() {
	dog := animals.Dog{}
	cat := animals.Cat{}

	fmt.Println(Perform(cat), Perform(dog))
}
