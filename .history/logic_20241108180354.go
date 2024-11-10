package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func (h Hero) attack(e *Enemy) {
	var weapon string
	fmt.Println("Choose your weapon")
	fmt.Scanln(&weapon)
	modifier := h.weapons[weapon]
	damage := rand.Intn(h.damage) + (1 + modifier)
	e.health -= damage
	if e.health <= 0 {
		e.dead = true
	}
	fmt.Printf("%s did %d points of damage, the %s has %d points of health left\n", h.name, damage, e.class, e.health)
}

func addItem[I Weapon | Equipment](i I) I {
	return i
}
func (h *Hero) getWeapon(w Weapon) {
	weapon := addItem(w)
	h.weapons[weapon.name] = weapon.damage
	fmt.Println(h.weapons)
}
func (h *Hero) getItem(s Equipment) {
	item := addItem(s)
	h.supplies[item.name] = item.amount
	fmt.Println(h.supplies)
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
		e.weakened = true
	case LIGHTNING:
		s.name = "Lightning"
		s.damage = 14
		e.repelled = true
	case FROST:
		s.name = "Frost"
		s.damage = 10
		e.slowed = true
	}
	e.health -= s.damage
	if e.health <= 0 {
		e.dead = true
	}
	fmt.Printf("The %s cast %s and did %d points of damage, %s has %d points of health left\n", h.name, s.name, s.damage, e.class, e.health)
}
func (h Hero) checkCondition() (b bool) {
	if h.weakened {
		fmt.Printf("%s is weakened\n", h.name)
		b = false
	}
	if h.slowed {
		fmt.Printf("%s is slowed\n", h.name)
		b = false
	}
	if h.imbued {
		fmt.Printf("%s is imbued\n", h.name)
		b = false
	}
	if h.dead {
		fmt.Printf("%s is dead\n", h.name)
		b = true
	}
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
		hP.weapons = map[string]int{"musket": 6}
		hP.supplies = map[string]int{}
		hP.slowed = false
		hP.weakened = false
		hP.imbued = false
		hP.dead = false
	case VAMPIRE:
		hP.class = "Vampire"
		hP.strength = 8
		hP.agility = 8
		hP.damage = 6
		hP.health = 4
		hP.weapons = map[string]int{"dagger": 4}
		hP.supplies = map[string]int{}
		hP.slowed = false
		hP.weakened = false
		hP.imbued = false
		hP.dead = false
	case WEREWOLF:
		hP.class = "Werewolf"
		hP.strength = 10
		hP.agility = 10
		hP.damage = 8
		hP.health = 10
		hP.weapons = map[string]int{"axe": 8}
		hP.supplies = map[string]int{}
		hP.slowed = false
		hP.weakened = false
		hP.imbued = false
		hP.dead = false
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
		eP.weapons = map[string]int{"cat of nine": 3}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false
	case HUNTER:
		eP.class = "Hunter"
		eP.strength = 8
		eP.agility = 8
		eP.damage = 6
		eP.health = 4
		eP.weapons = map[string]int{"bow": 5}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false
	case KNIGHT:
		eP.class = "Knight"
		eP.strength = 10
		eP.agility = 10
		eP.damage = 8
		eP.health = 10
		eP.weapons = map[string]int{"sword": 8}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false
	}

	return newEnemy, eP
}

func (e Enemy) attack(h *Hero) {
	damage := rand.Intn(e.damage) + 1
	h.health -= damage
	if h.health <= 0 {
		h.dead = true
	}
	fmt.Printf("The %s did %d points of damage, %s has %d points of health left\n", e.class, damage, h.name, h.health)
}

func (e Enemy) castSpell(h *Hero) {
	cl := rand.Intn(3) + 1
	var s = Spell{}

	switch cl {
	case FIREBALL:
		s.name = "Fireball"
		s.damage = 12
		h.weakened = true
	case LIGHTNING:
		s.name = "Lightning"
		s.damage = 14
	case FROST:
		s.name = "Frost"
		s.damage = 10
		h.slowed = true
	}
	h.health -= s.damage
	if h.health <= 0 {
		h.dead = true
	}
	fmt.Printf("The %s cast %s and did %d points of damage, %s has %d points of health left\n", e.class, s.name, s.damage, h.name, h.health)
}

func (h Hero) speak() {
	fmt.Println("My dark master will be pleased")
}
func (e Enemy) speak() {
	fmt.Println("I am safe in the light of the lord")
}
func (e Enemy) checkCondition() (b bool) {
	if e.weakened {
		fmt.Printf("%s is weakened\n", e.class)
		b = false
	}
	if e.slowed {
		fmt.Printf("%s is slowed\n", e.class)
		b = false
	}
	if e.repelled {
		fmt.Printf("%s is repelled\n", e.class)
		b = false
	}
	if e.dead {
		fmt.Printf("%s is dead\n", e.class)
		b = true
	}
	return b
}

type Character interface {
	speak()
	checkCondition() bool
}

func talk(c Character) {
	c.speak()

}

func checkConditions(c Character) {
	c.checkCondition()

}
func (h Hero) turn(e *Enemy) {
	h.attack(e)
	h.castSpell(e)
	talk(h)
}
func (e Enemy) turn(h *Hero) {
	e.attack(h)
	e.castSpell(h)
	talk(e)
}
func gameTurn(h *Hero, e *Enemy) {
	h.turn(e)
	e.turn(h)
}