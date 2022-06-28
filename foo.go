package pfoobar

import "math"

func Foo() string {
	return "FOO"
}

func Bar(s string) string {
	return s + "BAR"
}

//Return area of circle with radius r. Very hard task, of course.
func CircleArea(r float32) float32 {
	return math.Pi * r * r
}
