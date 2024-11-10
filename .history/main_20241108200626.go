package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var heroes []Hero
var villains []Enemy
var sword = Weapon{name: "sword", damage: 5}
var rope = Equipment{name: "rope", amount: 1}

func main() {
	rand.NewSource(time.Now().UnixNano())
	player1, p1 := spawnPlayer()
	enemy1, e1 := spawnEnemy()
	heroes = append(heroes, player1)
	villains = append(villains, enemy1)
	gameTurn(p1, e1)

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
