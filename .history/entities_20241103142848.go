package main

type Hero struct {
	name      string
	class     string
	strength  int
	agility   int
	health    int
	inventory []object
}
type Enemy struct {
	name     string
	strength int
	agility  int
	health   int
	x        float32
	y        float32
	height   float32
	width    float32
}
