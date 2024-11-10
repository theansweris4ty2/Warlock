package main

// TODO Should items be structs or maps
type Weapon struct {
	name   string
	damage int
}

var sword = Weapon{
	name:   "sword",
	damage: 6,
}
var bow = Weapon{
	name:   "bow",
	damage: 4,
}
var musket = Weapon{
	name:   "musket",
	damage: 5,
}
var axe = Weapon{
	name:   "axe",
	damage: 8,
}

type Armor struct {
	name    string
	defense int
}

var leather = Armor{
	name:    "leather armor",
	defense: 2,
}
var chain = Armor{
	name:    "chain mail",
	defense: 4,
}
var plate = Armor{
	name:    "plate armor",
	defense: 6,
}

type Potion struct {
	name      string
	condition string
}

var healing = Potion{
	name:      "healing potion",
	condition: "+2 health",
}
var speed = Potion{
	name:      "speed potion",
	condition: "+2 speed",
}
var mana = Potion{
	name:      "mana potion",
	condition: "+2 mana",
}
