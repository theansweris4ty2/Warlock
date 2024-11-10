package main

type Hero struct {
	name     string
	class    string
	strength int
	agility  int
	health   int
	damage   int
	// armory   []Weapon
	// x        float32
	// y        float32
	// height   float32
	// width    float32
}
type Enemy struct {
	class    string
	strength int
	agility  int
	health   int
	damage   int
	// x        float32
	// y        float32
	// height   float32
	// width    float32
}

// type Weapon struct {
// 	name   string
// 	damage int
// }

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
