package main

type Spell struct {
	name    string
	effects string
	damage  int
}

const (
	SLOW = iota + 1
	FOG
	FLY
)
