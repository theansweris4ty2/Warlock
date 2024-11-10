package main

// TODO Should items be structs or maps
type Weapon struct {
	name   string
	damage int
}

type Equipment struct {
	name   string
	amount int
}

type Armor struct {
	name    string
	defense int
}

type Potion struct {
	name      string
	condition string
	amount    int
}
