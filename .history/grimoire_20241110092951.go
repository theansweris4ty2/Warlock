package main

type Spell struct {
	name   string
	damage int
	// TODO decide whether there should be a condition field
}

const (
	FIREBALL = iota + 1
	LIGHTNING
	FROST
)
