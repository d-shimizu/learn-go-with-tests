package clockface

import (
	"math"
	"time"
)

func secondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * (math.Pi / 30)
}
