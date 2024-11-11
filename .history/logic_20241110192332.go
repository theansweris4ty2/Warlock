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
	modifier := h.weapons[weapon].damage
	damage := rand.Intn(h.damage) + (1 + modifier)
	if !e.dead {
		e.health -= damage
	}
	if e.health <= 0 {
		e.dead = true
	}
	fmt.Printf("%s did %d points of damage, the %s has %d points of health left\n", h.name, damage, e.class, e.health)
}

func addItem[I Weapon | Potion | Armor](i I) I {
	return i
}
func (h *Hero) getWeapon(w Weapon) {
	weapon := addItem(w)
	h.weapons[weapon.name] = w
	fmt.Println(h.weapons)
}
func (h *Hero) getArmor(a Armor) {
	item := addItem(a)
	h.armor[item.name] = a
	fmt.Println(h.armor)
}
func (h *Hero) getPotion(p Potion) {
	item := addItem(p)
	h.potions[item.name] = p
	fmt.Println(h.potions)
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
func (h Hero) checkCondition() {
	// TODO add logic so that conditions impact player
	if h.weakened {
		fmt.Printf("%s is weakened\n", h.name)
	}
	if h.slowed {
		fmt.Printf("%s is slowed\n", h.name)
	}
	if h.imbued {
		fmt.Printf("%s is imbued\n", h.name)
	}
	if h.dead {
		fmt.Printf("%s is dead\n", h.name)
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
	case KNIGHT:
		hP.class = "Knight"
		hP.strength = 6
		hP.agility = 7
		hP.damage = 4
		hP.health = 30
		hP.sprite = Sprite{}
		hP.weapons = map[string]Weapon{}
		hP.potions = map[string]Potion{}
		hP.armor = map[string]Armor{}
		hP.slowed = false
		hP.weakened = false
		hP.imbued = false
		hP.dead = false
	case ROGUE:
		hP.class = "Rogue"
		hP.strength = 8
		hP.agility = 8
		hP.damage = 6
		hP.health = 4
		hP.sprite = Sprite{}
		hP.weapons["spear"] = Weapon{name: "spear", damage: 6}
		hP.potions = map[string]Potion{}
		hP.slowed = false
		hP.weakened = false
		hP.imbued = false
		hP.dead = false
	case WIZARD:
		hP.class = "Wizard"
		hP.strength = 10
		hP.agility = 10
		hP.damage = 8
		hP.health = 10
		hP.sprite = Sprite{}
		hP.weapons["axe"] = Weapon{name: "axe", damage: 8}
		hP.potions = map[string]Potion{}
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
	case ORC_ROGUE:
		eP.class = "Orc Rogue"
		eP.damage = 4
		eP.health = 30
		eP.agility = 4
		eP.sprite = Sprite{}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false
	case ORC_SHAMAN:
		eP.class = "Orc Shaman"
		eP.damage = 6
		eP.health = 30
		eP.agility = 6
		eP.sprite = Sprite{}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false
	case ORC_WARRIOR:
		eP.class = "Orc Warrior"
		eP.damage = 8
		eP.health = 30
		eP.agility = 3
		eP.sprite = Sprite{}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false
	case SKELETON_ROGUE:
		eP.class = "Skeleton Rogue"
		eP.damage = 4
		eP.health = 30
		eP.agility = 4
		eP.sprite = Sprite{}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false
	case SKELETON_MAGE:
		eP.class = "Skeleton Mage"
		eP.damage = 6
		eP.health = 30
		eP.agility = 6
		eP.sprite = Sprite{}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false
	case SKELETON_WARRIOR:
		eP.class = "Skeleton Warrior"
		eP.damage = 8
		eP.health = 30
		eP.agility = 3
		eP.sprite = Sprite{}
		eP.slowed = false
		eP.weakened = false
		eP.repelled = false
		eP.dead = false

	}

	return newEnemy, eP
}

func (e Enemy) attack(h *Hero) {
	damage := rand.Intn(e.damage) + 1
	if !h.dead {
		h.health -= damage
	}
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
func (e Enemy) checkCondition() {
	if e.weakened {
		fmt.Printf("%s is weakened\n", e.class)
	}
	if e.slowed {
		fmt.Printf("%s is slowed\n", e.class)
	}
	if e.repelled {
		fmt.Printf("%s is repelled\n", e.class)
	}
	if e.dead {
		fmt.Printf("%s is dead\n", e.class)
	}
}

func (h Hero) turn(e *Enemy) {
	// FIXME - change logic so it does not require passing enemy parameter, player should choose enemy through targeting
	h.checkCondition()
	h.attack(e)
	h.castSpell(e)
	h.speak()
}
func (e Enemy) turn(h *Hero) {
	// FIXME - change logic - enemy will always attack player 1
	e.checkCondition()
	e.attack(h)
	e.castSpell(h)
	e.speak()
}
