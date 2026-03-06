package main

import "fmt"

type Item struct {
	Quantity int
	Price    float64
}

func main() {
	inventory := map[string]Item{
		"Apple":   {Quantity: 10, Price: 1.50},
		"Bananas": {Quantity: 8, Price: 0.75},
		"Mango":   {Quantity: 4, Price: 2.80},
		"Grapes":  {Quantity: 6, Price: 3.20},
	}

	AddItems(inventory, "Apple", 5)
	AddItems(inventory, "Bananas", 2)

	UpdatePrice(inventory, "Mango", 3.10)

	fmt.Println("Items with price > $2.00:")
	for name, item := range inventory {
		if item.Price > 2.00 {
			fmt.Printf("- %s: qty=%d, price=$%.2f\n", name, item.Quantity, item.Price)
		}
	}

	delete(inventory, "Bananas")

	var totalValue float64
	for _, item := range inventory {
		totalValue += float64(item.Quantity) * item.Price
	}
	fmt.Printf("Total inventory value: $%.2f\n", totalValue)
}

func AddItems(inventory map[string]Item, item string, quantity int) {
	value, exists := inventory[item]
	if !exists {
		inventory[item] = Item{Quantity: quantity, Price: 0}
		return
	}

	value.Quantity += quantity
	inventory[item] = value
}

func UpdatePrice(inventory map[string]Item, item string, price float64) {
	value, exists := inventory[item]
	if !exists {
		return
	}

	value.Price = price
	inventory[item] = value
}
