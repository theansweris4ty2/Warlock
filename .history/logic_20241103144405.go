package main

import (
	"math/rand"
)

func (h *Hero) attack(e *Enemy) {
	damage := rand.Intn(h.damage)
	e.health -= damage
}

func createPlayer() (Hero, *Hero) {
	var hP *Hero = newPlayer Hero{}
	hP = &newPlayer
	return newPlayer, hP
}
