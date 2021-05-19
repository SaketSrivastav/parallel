package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
	borderX  float64
	borderY  float64
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)
	if next.x >= b.borderX || next.x < 0 {
		b.velocity = Vector2D{-b.velocity.x, b.velocity.y} // bounce effect when we reach border of screen
	}

	if next.y >= b.borderY || next.y < 0 {
		b.velocity = Vector2D{b.velocity.x, -b.velocity.y} // bounce effect when we reach border of screen
	}
}

func (b *Boid) Start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func CreateBoid(id int, xLimit float64, yLimit float64) *Boid {
	b := &Boid{
		position: Vector2D{x: rand.Float64() * xLimit, y: rand.Float64() * yLimit},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1, y: (rand.Float64() * 2) - 1},
		id:       id,
		borderX:  xLimit,
		borderY:  yLimit,
	}

	go b.Start()
	return b
}
