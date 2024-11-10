package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var characters []Character

func main() {
	rand.NewSource(time.Now().UnixNano())
	player1, p1 := spawnPlayer()
	enemy1, e1 := spawnEnemy()
	characters = append(characters, player1)
	characters = append(characters, enemy1)
	p1.attack(e1)
	e1.attack(p1)
	p1.castSpell(e1)
	e1.castSpell(p1)
	fmt.Println(p1)
	fmt.Println(e1)
	talk(p1)
	checkConditions(p1)

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
