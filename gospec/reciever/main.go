package main

import "fmt"

type Person struct {
	name string
}

func (p Person) Set(name string) {
	p.name = name
}

func (p Person) Greet() {
	fmt.Printf("Hello, I'm %s\n", p.name)
}

type PPerson struct {
	name string
}

func (p *PPerson) Set(name string) {
	p.name = name
}

func (p *PPerson) Greet() {
	fmt.Printf("Hello, I'm %s\n", p.name)
}

func main() {
	p := &Person{name: "tMinamiii"}
	p.Greet()
	p.Set("minami")
	p.Greet()

	pp := &PPerson{name: "tMinamiii"}
	pp.Greet()
	pp.Set("minami")
	pp.Greet()
}
