package wildnature

import "fmt"

//Animal represents a struct with basic animal attributes and actions
type animal struct {
	species string
	age     int
	color   string
}

//Cat implements animal cat
type Cat struct {
	animal
	name    string
	hostile bool
}

//Dog implements animal dog
type Dog struct {
	animal
	name     string
	friendly bool
}

//SoundMaker implements an entity which sounds
type SoundMaker interface {
	MakeSound()
}

//GetAge returns animal's age
func (a *animal) GetAge() int {
	return a.age
}

//GetName returns animal's name
func (c *Cat) GetName() string {
	return c.name
}

//GetName returns animal's name
func (d *Dog) GetName() string {
	return d.name
}

//New returns a new instance of a Cat type
func (c *Cat) New() *Cat {
	newcat := Cat{
		animal: animal{
			species: "Felis catus",
			age:     5,
			color:   "black",
		},
		name:    "Tom",
		hostile: true,
	}
	return &newcat
}

//New returns a new instance of a Dog type
func (d *Dog) New() *Dog {
	newdog := Dog{
		animal: animal{
			species: "Felis catus",
			age:     7,
			color:   "black",
		},
		name:     "Charly",
		friendly: true,
	}
	return &newdog
}

// MakeSound demonstrates cat's voice
func (c *Cat) MakeSound() {
	if c.hostile {
		fmt.Println("MEOW!!!!!!")
	} else {
		fmt.Println("mew")
	}
}

// MakeSound demonstrates dog's voice
func (d *Dog) MakeSound() {
	if !d.friendly {
		fmt.Println("WWUFF!!!!!!")
	} else {
		fmt.Println("wof")
	}
}
