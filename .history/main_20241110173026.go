package main

import (
	"ebiten"
	"math/rand"
	"time"
)

func main() {
	heroes := make(map[string]Hero)
	villains := make(map[string]Enemy)
	rand.NewSource(time.Now().UnixNano())
	player1, p1 := spawnPlayer()
	player2, p2 := spawnPlayer()
	enemy1, e1 := spawnEnemy()
	heroes[p1.name] = player1
	heroes[p2.name] = player2
	villains[e1.class] = enemy1
	p1.getWeapon(Weapon{name: "musket", damage: 6})
	p1.getPotion(Potion{name: "Healing", condition: "heal 6"})
	p1.getArmor(Armor{name: "leather", defense: 2})
	for len(heroes) > 0 && len(villains) > 0 {
		for _, p := range heroes {
			p.turn(e1)
			if e1.dead {
				delete(villains, e1.class)
			}
		}
		for _, e := range villains {
			e.turn(p1)
			if p1.dead {
				delete(heroes, p1.name)
			}
		}
	}

}
