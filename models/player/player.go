package player

import "strangeadventure/models/items"


type Person struct {
	Name string
	World int
	Level int
	Health int
	Rank string
	Backpack
}

// Геттеры
func (p *Person) GetInfo() Person {
	return *p
}
func (p *Person) GetName() string {
	return p.Name
}

// Сеттеры 
func (p *Person) SetName(name string) {
	p.Name = name
}


type Player interface {

	//This func return player info
    //name - GetInfo().Name
	// world - GetInfo().World
	// level - GetInfo().Level
	// health - GetInfo().Health
	// rank - GetInfo().Rank
	// backpack - GetInfo().Backpack
	// and more...
	GetInfo() Person

	// This func return player name
	GetName() string

	// This func take name and change palyer name
	SetName(string)
	
	/*
	Pass [name,rarity,types,price] of new item
	This func add item in Backpack.Invetory 
	*/
	Push(string,string,string,int)
	
	// This function delete item with pass id
	Delete(int)
}

// Pass player name as arg
func PersonConstructor(name string) Player {
	p := Person{
		Name: name,
		Level: 0,
		Health: 100,
		World: 1,
		Rank: "новичок",
		Backpack: Backpack{
			Size: 10,
			Inventory: []items.Item{
				{Id: 1,Name: "Цветочек",Price: 0,Rarity: "обычный",Types: "предмет"},
			},
		},
	}
	return &p
}



