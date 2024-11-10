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
	var hP *Hero = &newPlayer
	return newPlayer, hP
}

func (e *Enemy) attack(h *Hero) {
	damage := rand.Intn(e.damage)
	h.health -= damage
}
