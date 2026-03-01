package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

type Contact struct {
	Name    string
	Number  string
	Address Address
}

type Address struct {
	Street string
	City   string
}

func main() {
	p := Person{
		Name:  "Aakash",
		Age:   22,
		Email: "aakash19here@gmail.com",
	}

	c := struct {
		Name string
		Type string
	}{
		Name: "Toyota Camry",
		Type: "Sedan",
	}

	fmt.Printf("The name is %s , age is %d and email is %s \n", p.Name, p.Age, p.Email)
	fmt.Printf("The Car is %s , type is %s\n", c.Name, c.Type)

}
