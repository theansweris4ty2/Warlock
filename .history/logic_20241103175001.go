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

func spawnPlayer() (Hero, *Hero) {
	var cl string
	var name string
	var newPlayer = Hero{}
	var hP *Hero = &newPlayer
	fmt.Print("What is your name hero?\n")
	fmt.Scanln(&name)
	hP.name = name
	fmt.Print("Choose your class: 1. for Warlock, 2. Vampire, or 3. for Werewolf\n")
	fmt.Scanln(&cl)
	choice, _ := strconv.Atoi(strings.Trim(cl, "\n"))
	switch choice {
	case WARLOCK:
		hP.class = "Warlock"
		hP.strength = 6
		hP.agility = 7
		hP.damage = 4
		hP.health = 6
	case VAMPIRE:
		hP.class = "Vampire"
		hP.strength = 8
		hP.agility = 8
		hP.damage = 6
		hP.health = 4
	case WEREWOLF:
		hP.class = "Werewolf"
		hP.strength = 10
		hP.agility = 10
		hP.damage = 8
		hP.health = 10
	}

	return newPlayer, hP
}

func spawnEnemy() (Enemy, *Enemy) {
	var cl string
	var name string
	var newEnemy = Enemy{}
	var eP *Enemy = &newEnemy
	fmt.Print("What is your name hero?\n")
	fmt.Scanln(&name)
	eP.name = name
	fmt.Print("Choose your class: 1. for Warlock, 2. Vampire, or 3. for Werewolf\n")
	fmt.Scanln(&cl)
	choice, _ := strconv.Atoi(strings.Trim(cl, "\n"))
	switch choice {
	case WARLOCK:
		eP.class = "Warlock"
		eP.strength = 6
		eP.agility = 7
		eP.damage = 4
		eP.health = 6
	case VAMPIRE:
		eP.class = "Vampire"
		eP.strength = 8
		eP.agility = 8
		eP.damage = 6
		eP.health = 4
	case WEREWOLF:
		eP.class = "Werewolf"
		eP.strength = 10
		eP.agility = 10
		eP.damage = 8
		eP.health = 10
	}

	return newEnemy, eP
}

func (e *Enemy) attack(h *Hero) {
	damage := rand.Intn(e.damage)
	h.health -= damage
	fmt.Printf("%s did %d points of damage, %s has %d points of health left", e.name, damage, h.name, h.health)
}
