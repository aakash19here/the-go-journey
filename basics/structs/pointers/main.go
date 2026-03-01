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

	// reference to the variable / address of person
	modifyName(&p)

	fmt.Println("Name after", p.Name)

}

// pointer to person
func modifyName(p *Person) {
	p.Name = "Akku"
	fmt.Println("Name inside scope", p.Name)
}
