package main

import (
	"image"
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "math/rand"
	// "time"
)

type Game struct {
	PlayerImage *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{150, 175, 255, 255})
	screen.DrawImage(g.PlayerImage.SubImage(image.Rect(0, 0, 60, 60)).(*ebiten.Image), &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	playerImage, _, err := ebitenutil.NewImageFromFile("assets/images/Heroes/Knight/Run/Run-Sheet.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{PlayerImage: playerImage}); err != nil {
		log.Fatal(err)
	}
}
