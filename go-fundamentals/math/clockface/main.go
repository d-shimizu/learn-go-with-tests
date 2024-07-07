package clockface

import "time"

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}
	p = Point{p.X, -p.Y}
	p = Point{p.X + 150, p.Y + 150}
	return p
}
