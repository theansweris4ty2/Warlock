package main

type Spell struct {
	name    string
	effects string
	damage  int
}

const (
	FIREBALL = iota + 1
	LIGHTNING
	FROST
)
