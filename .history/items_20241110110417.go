package main

// TODO Should items be structs or maps
type Weapon struct {
	name   string
	damage int
}

var sword Weapon

type Armor struct {
	name    string
	defense int
}

type Potion struct {
	name      string
	condition string
}
