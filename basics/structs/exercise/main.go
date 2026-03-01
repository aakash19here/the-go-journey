package main

import (
	"fmt"
	"slices"
)

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

func main() {
	p1 := Player{
		Name: "fallen_god",
		Inventory: []Item{
			{
				Name: "Axe",
				Type: "Weapon",
			},
			{
				Name: "Mk4",
				Type: "Gun",
			},
		},
	}

	//PickUpItem
	p1.PickUpItem(Item{
		Name: "Pan",
		Type: "Melee",
	})

	printItems(p1.Inventory)

	delimiter()

	p1.DropItem("Mk4")

	printItems(p1.Inventory)

}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}

func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			// p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			p.Inventory = slices.Delete(p.Inventory, i, i+1)
		}
	}

}

func printItems(items []Item) {
	for _, item := range items {
		fmt.Println("Inventory item is", item.Name, "and the type is", item.Type)
	}
}

func delimiter() {
	fmt.Println("===============")
}
