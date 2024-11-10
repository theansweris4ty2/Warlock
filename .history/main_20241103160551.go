package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var players []Hero
var enemies []Enemy

func main() {
	rand.NewSource(time.Now().UnixNano())
	player1, p1 := createPlayer()
	enemy1, e1 := createEnemy()
	players = append(players, player1)
	enemies = append(enemies, enemy1)
	p1.attack(e1)
	e1.attack(p1)
	fmt.Println(p1)
	// fmt.Println(p1.class)

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
