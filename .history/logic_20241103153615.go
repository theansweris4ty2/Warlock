package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func (h *Hero) attack(e *Enemy) {
	damage := rand.Intn(h.damage)
	e.health -= damage
	fmt.Printf("%s did %d points of damage, %s has %d points of health left", h.name, damage, e.name, e.health)
}

func createPlayer() (Hero, *Hero) {
	var cl string
	var newPlayer = Hero{}
	var hP *Hero = &newPlayer
	fmt.Println("Choose your class: 1. for Warlock, 2. Vampire, or 3. for Werewolf")
	fmt.Scanf("%s", &cl)
	choice, _ := strconv.Atoi(strings.Trim(cl, "\n"))
	switch choice {
	case WARLOCK:
		hP.strength = 6
		hP.agility = 7
		hP.health = 6
	case VAMPIRE:
		hP.strength = 8
		hP.agility = 8
		hP.health = 4
	case WEREWOLF:
		hP.strength = 10
		hP.agility = 10
		hP.health = 10
	}

	return newPlayer, hP
}

func (e *Enemy) attack(h *Hero) {
	damage := rand.Intn(e.damage)
	h.health -= damage
	fmt.Printf("%s did %d points of damage, %s has %d points of health left", e.name, damage, h.name, h.health)
}
