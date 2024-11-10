package main

type Weapon struct {
	name   string
	damage int
}

type Equipment struct {
	name   string
	amount int
}
type Item interface {
	Weapon | Equipment
}
