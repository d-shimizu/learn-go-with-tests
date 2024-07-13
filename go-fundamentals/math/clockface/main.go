package main

import (
	"os"
	"time"

	clockface "github.com/d-shimizu/learn-go-with-tests/go-fundamentals/math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
