package player

import (
	"fmt"
	"strangeadventure/models/items"
)

type Backpack struct {
	Length int
	Size int
	Inventory []items.Item
}
func (b *Backpack) Push(name,rarity,types string, price int) {
	if b.Length < b.Size {
		b.Inventory = append(b.Inventory, items.Item{
			Id: b.Inventory[len(b.Inventory)-1].Id + 1,
			Name: name,
			Rarity: rarity,
			Types: types,
			Price: price,
		})
		b.Length = len(b.Inventory)
	} else {
		fmt.Println("Ваш рюкзак переполнен")
	}
}
func (b *Backpack) Delete(elementId int) {
	for i, i2 := range b.Inventory {
		if i2.Id == elementId {
			b.Inventory = append(b.Inventory[:i],b.Inventory[i+1:]...)
		}
	}
}