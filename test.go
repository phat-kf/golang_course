package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak()
}

// Implement the interface in a struct
type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "says Woof!")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "says Meow!")
}

func main() {
	i := 1
	for i <= 5 {
		if i == 3 {
			goto skip // Jumps to the label
		}
		fmt.Println(i)
		i++
	}
skip:
	fmt.Println("Skipped part of the loop")
}
