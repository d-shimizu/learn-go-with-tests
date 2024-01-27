package main

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// 長方形の周囲を計算する関数
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// 長方形の面積を計算する関数
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
