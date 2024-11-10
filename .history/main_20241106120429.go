package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var characters []Character
var sword = Weapon{name: "sword", damage: 5}

func main() {
	rand.NewSource(time.Now().UnixNano())
	_, p1 := spawnPlayer()
	_, e1 := spawnEnemy()
	characters = append(characters, p1)
	characters = append(characters, e1)
	p1.attack(e1)
	e1.attack(p1)
	p1.castSpell(e1)
	e1.castSpell(p1)
	talk(p1)
	for _, ch := range characters {
		checkConditions(ch)
	}
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
