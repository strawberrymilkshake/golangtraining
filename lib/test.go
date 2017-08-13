package main

import (
	"fmt"

	"github.com/strawberrymilkshake/golangtraining/lib/wildnature"
)

func sound(s ...wildnature.SoundMaker) {
	for _, v := range s {
		v.MakeSound()
	}
}

func main() {
	var c *wildnature.Cat
	c = c.New()

	var d *wildnature.Dog
	d = d.New()

	sound(c, d)
	fmt.Println("Cat's name is", c.GetName(), "and he is", c.GetAge(), "years old")
	fmt.Println("Dog's name is", d.GetName(), "and he is", d.GetAge(), "years old")
}
