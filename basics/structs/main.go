package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	p := Person{
		Name:  "Aakash",
		Age:   22,
		Email: "aakash19here@gmail.com",
	}

	fmt.Printf("The name is %s , age is %d and email is %s \n", p.Name, p.Age, p.Email)

}
