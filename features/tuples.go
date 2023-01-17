package features

import "math"

const EPSILON = 0.00001

// type Point struct {
// 	x float64
// 	y float64
// 	z float64
// }

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
	// data [4]interface{}
}

// Point returns a point
func Point(x, y, z float64) Tuple {
	w := 1.0
	t := Tuple{x, y, z, w}
	return t
}

// Vector returns a vector
func Vector(x, y, z float64) Tuple {
	w := 0.0
	t := Tuple{x, y, z, w}
	return t
}

func Add(t1, t2 Tuple) Tuple {
	t := Tuple{}
	t.x = t1.x + t2.x
	t.y = t1.y + t2.y
	t.z = t1.z + t2.z
	t.w = t1.w + t2.w
	// TODO: fix this
	if t.w > 1 {
		panic("O fuk i abded")
	}
	return t
}

func Subtract(t1, t2 Tuple) Tuple {
	t := Tuple{}
	t.x = t1.x - t2.x
	t.y = t1.y - t2.y
	t.z = t1.z - t2.z
	t.w = t1.w - t2.w
	// TODO: fix this
	if t.w < 0 {
		panic("O fuk i subtarcted")

	}
	return t
}

func IsPoint(t Tuple) bool {
	return t.w == 1.0
}

func IsVector(t Tuple) bool {
	return t.w == 0.0
}

func Equal(t1, t2 Tuple) bool {
	return equal(t1.x, t2.x) &&
		equal(t1.y, t2.y) &&
		equal(t1.z, t2.z) &&
		equal(t1.w, t2.x)
}

func equal(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

// Point returns a point
// func Point(x, y, z float64) tuple {
// 	t := tuple{[4]interface{}{x, y, z, 1.0}}
// 	return t
// }

// // Vector returns a vector
// func Vector(x, y, z float64) tuple {
// 	t := tuple{[4]interface{}{x, y, z, 0}}
// 	return t
// }

// func isPoint(t tuple) bool {
// 	return t.data[3] == 1.0
// }

// func isVector(t tuple) bool {
// 	return t.data[3] == 0
// }
