package main

import (
	"fmt"
	f "ray-tracing/features"
)

func main() {
	p := f.Projectile{
		Position: f.Point(0.0, 1.0, 0.0),
		Velocity: f.Normalize(f.Vector(1.0, 1.0, 0.0)),
	}

	e := f.Environment{
		Gravity: f.Vector(0.0, -0.1, 0.0),
		Wind:    f.Vector(-0.01, 0.0, 0.0),
	}

	tickCount := 0
	for p.Position.Y > 0 {
		fmt.Println(p.Position)
		p = f.Tick(e, p)
		tickCount += 1
		fmt.Printf("There's been %d ticks\n", tickCount)
	}
	// fmt.Println(p)

}
