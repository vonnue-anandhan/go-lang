package methodsets

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	d.name = "Rover"
	fmt.Println("My name is: ", d.name, " and I'm walking")
}

func (d *dog) run() {
	d.name = "Rover"
	fmt.Println("My name is: ", d.name, " and I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func TestSets() {
	dog1 := dog{name: "Henry"}

	dog1.walk()
	fmt.Println("New dog1 name: ", dog1.name) // Henry
	dog1.run()
	fmt.Println("New dog1 name: ", dog1.name) //  Rover

	// youngRun(dog1) ERROR!!

	dog2 := &dog{name: "Padget"}

	dog2.walk()
	fmt.Println("New dog2 name: ", dog2.name) // Padget
	dog2.run()
	fmt.Println("New dog2 name: ", dog2.name) // Rover

	youngRun(dog2)
}
