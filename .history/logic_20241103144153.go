package main

import (
	"math/rand"
)

func (h *Hero) attack(e *Enemy) {
	damage := rand.Intn(h.damage)
	e.health -= damage
}

func createPlayer() (Hero, *Hero) {
	var newPlayer = Hero{}
	var hP *Hero
	hP = &newPlayer
	return newPlayer, hP
}
