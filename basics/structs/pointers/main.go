package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	p := Person{
		Name: "Aakash",
	}

	fmt.Println("Name before", p.Name)

	p.modifyName()

	fmt.Println("Name after", p.Name)

}

// method on person struct or called as method receiver
// pointer to person
func (p *Person) modifyName() {
	p.Name = "Akku"
	fmt.Println("Name inside scope", p.Name)
}
