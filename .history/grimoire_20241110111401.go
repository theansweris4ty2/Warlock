package main

type Spell struct {
	name      string
	damage    int
	condition string
}

const (
	FIREBALL = iota + 1
	LIGHTNING
	FROST
)
