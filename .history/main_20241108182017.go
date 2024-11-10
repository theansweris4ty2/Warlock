package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var characters []Character
var sword = Weapon{name: "sword", damage: 5}
var rope = Equipment{name: "rope", amount: 1}

func main() {
	rand.NewSource(time.Now().UnixNano())
	_, p1 := spawnPlayer()
	_, e1 := spawnEnemy()
	characters = append(characters, p1)
	characters = append(characters, e1)
	gameTurn(p1, e1)
	for i, ch := range characters {
		b := checkConditions(ch)
		if b {
			characters = append(characters[:i], characters[i:]...)
		}

	}
	p1.getWeapon(sword)
	p1.getItem(rope)
	for _, ch := range characters {
		fmt.Print(ch)
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
