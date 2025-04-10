package main

import (
	"fmt"

	"github.com/AaronKluck/goplay/play"
)

func main() {
	fmt.Println("Hello, World!")
	play.SelectorChan(
		play.NewPerson,
		play.NewCar,
		func(val play.Stringable) {
			fmt.Println("Received 1:", val.Repr())
		},
		func(val play.Stringable) {
			fmt.Println("Received 2:", val.Repr())
		},
	)
}
