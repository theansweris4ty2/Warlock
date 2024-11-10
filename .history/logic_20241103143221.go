package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)
rand.NewSource(time.Now().UnixNano())
	num := rand.Intn(20)
	fmt.Printf("%d", num)