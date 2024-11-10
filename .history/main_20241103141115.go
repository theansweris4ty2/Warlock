package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	num := rand.Intn(20)
	fmt.Printf("%d", num)
}
