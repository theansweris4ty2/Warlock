package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	num := rand.Intn(20)
	fmt.Printf("%d", num)
}
