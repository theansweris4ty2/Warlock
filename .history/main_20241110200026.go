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
	x, y        float64
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{150, 175, 255, 255})
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.x, g.y)
	screen.DrawImage(g.PlayerImage.SubImage(image.Rect(0, 0, 60, 65)).(*ebiten.Image), &ebiten.DrawImageOptions{})
	screen.DrawImage(g.PlayerImage.SubImage(image.Rect(61, 0, 121, 65)).(*ebiten.Image), &opts)
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

	if err := ebiten.RunGame(&Game{PlayerImage: playerImage, x: 100, y: 100}); err != nil {
		log.Fatal(err)
	}
}
