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
	switch(choice)
	return newPlayer, hP
}

func (e *Enemy) attack(h *Hero) {
	damage := rand.Intn(e.damage)
	h.health -= damage
	fmt.Printf("%s did %d points of damage, %s has %d points of health left", e.name, damage, h.name, h.health)
}