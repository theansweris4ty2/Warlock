package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func (h Hero) attack(e *Enemy) {
	damage := rand.Intn(h.damage)
	e.health -= damage
	fmt.Printf("%s did %d points of damage, the %s has %d points of health left\n", h.name, damage, e.class, e.health)
}
func (h Hero) castSpell(e *Enemy) {
	var cl string
	var s = Spell{}
	fmt.Print("Choose your spell: 1. for Fireball, 2. Lightning, or 3. for Frost\n")
	fmt.Scanln(&cl)
	choice, _ := strconv.Atoi(strings.Trim(cl, "\n"))
	switch choice {
	case FIREBALL:
		s.name = "Fireball"
		s.damage = 12
	case LIGHTNING:
		s.name = "Lightning"
		s.damage = 14
	case FROST:
		s.name = "Frost"
		s.damage = 10
	}
	e.health -= s.damage
	fmt.Printf("The %s cast %s and did %d points of damage, %s has %d points of health left\n", h.name, s.name, s.damage, e.class, e.health)
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
		hP.slowed = false
		hP.weakened = false
		hP.imbued = false
	case VAMPIRE:
		hP.class = "Vampire"
		hP.strength = 8
		hP.agility = 8
		hP.damage = 6
		hP.health = 4
		hP.slowed = false
		hP.weakened = false
		hP.weakened = false
	case WEREWOLF:
		hP.class = "Werewolf"
		hP.strength = 10
		hP.agility = 10
		hP.damage = 8
		hP.health = 10
		hP.slowed = false
		hP.weakened = false
		hP.imbued = false
	}

	return newPlayer, hP
}

func spawnEnemy() (Enemy, *Enemy) {
	spawned := rand.Intn(3) + 1
	var newEnemy = Enemy{}
	var eP *Enemy = &newEnemy
	switch spawned {
	case INQUISITOR:
		eP.class = "Inquisitor"
		eP.strength = 6
		eP.agility = 7
		eP.damage = 4
		eP.health = 6
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
	case HUNTER:
		eP.class = "Hunter"
		eP.strength = 8
		eP.agility = 8
		eP.damage = 6
		eP.health = 4
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
	case KNIGHT:
		eP.class = "Knight"
		eP.strength = 10
		eP.agility = 10
		eP.damage = 8
		eP.health = 10
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
	}

	return newEnemy, eP
}

func (e Enemy) attack(h *Hero) {
	damage := rand.Intn(e.damage)
	h.health -= damage
	fmt.Printf("The %s did %d points of damage, %s has %d points of health left\n", e.class, damage, h.name, h.health)
}

func (e Enemy) castSpell(h *Hero) {
	cl := rand.Intn(3) + 1
	var s = Spell{}

	switch cl {
	case FIREBALL:
		s.name = "Fireball"
		s.damage = 12
	case LIGHTNING:
		s.name = "Lightning"
		s.damage = 14
	case FROST:
		s.name = "Frost"
		s.damage = 10
	}
	h.health -= s.damage
	fmt.Printf("The %s cast %s and did %d points of damage, %s has %d points of health left\n", e.class, s.name, s.damage, h.name, h.health)
}

func (h Hero) threaten() {
	fmt.Println("I will swallow your soul")
}
func (e Enemy) threaten() {
	fmt.Println("Face the light and fear it's power.")

}

func (h Hero) speak() {
	fmt.Println("My dark master will be pleased")
}
func (e Enemy) speak() {
	fmt.Println("I am safe in the light of the lord")
}

type Character interface {
	threaten()
}
