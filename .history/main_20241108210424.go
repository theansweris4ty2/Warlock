package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var sword = Weapon{name: "sword", damage: 5}
var rope = Equipment{name: "rope", amount: 1}

func main() {
	heroes := make(map[string]Hero)
	villains := make(map[string]Enemy)
	rand.NewSource(time.Now().UnixNano())
	player1, p1 := spawnPlayer()
	enemy1, e1 := spawnEnemy()
	heroes[p1.name] = player1
	villains[e1.class] = enemy1
	for len(heroes) > 0 {
	}

	p1.getWeapon(sword)
	p1.getItem(rope)
}

func startGame() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
