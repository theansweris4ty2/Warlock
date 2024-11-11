package main

type sprite struct {
	img string
	x   float32
	y   float32
}
type Hero struct {
	name     string
	class    string
	strength int
	agility  int
	health   int
	damage   int
	sprite   sprite
	weapons  map[string]Weapon
	armor    map[string]Armor
	potions  map[string]Potion
	slowed   bool
	weakened bool
	imbued   bool
	dead     bool
	// height   float32
	// width    float32
}
type Enemy struct {
	class    string
	health   int
	damage   int
	sprite   sprite
	slowed   bool
	weakened bool
	repelled bool
	dead     bool
	// height   float32
	// width    float32
}

const (
	WARLOCK = iota + 1
	VAMPIRE
	WEREWOLF
)
const (
	INQUISITOR = iota + 1
	HUNTER
	KNIGHT
)
