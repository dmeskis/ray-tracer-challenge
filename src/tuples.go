package features

import (
	"math"
)

const EPSILON = 0.00001

// type Point struct {
// 	x float64
// 	y float64
// 	z float64
// }

// Might need to refactor Tuple to point/vector...
type Projectile struct {
	position Tuple // Point
	velocity Tuple // Vector
}

type Environment struct {
	gravity Tuple // Vector
	wind    Tuple // Vector
}

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
		equal(t1.w, t2.w)
}

func equal(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func Negate(t Tuple) Tuple {
	zero := Tuple{}
	return Subtract(zero, t)
}

func Multiply(t1 Tuple, scalar float64) Tuple {
	t := Tuple{}
	t.x = t1.x * scalar
	t.y = t1.y * scalar
	t.z = t1.z * scalar
	t.w = t1.w * scalar
	return t
}

func Divide(t1 Tuple, scalar float64) Tuple {
	t := Tuple{}
	t.x = t1.x / scalar
	t.y = t1.y / scalar
	t.z = t1.z / scalar
	t.w = t1.w / scalar
	return t
}

func Magnitude(t Tuple) float64 {
	radicand := (t.x * t.x) + (t.y * t.y) + (t.z * t.z) + (t.w * t.w)
	return math.Sqrt(radicand)
}

func Normalize(t1 Tuple) Tuple {
	m := Magnitude(t1)
	t := Tuple{}
	t.x = t1.x / m
	t.y = t1.y / m
	t.z = t1.z / m
	t.w = t1.w / m
	return t
}

func Dot(t1, t2 Tuple) (scalar float64) {
	scalar = (t1.x * t2.x) +
		(t1.y * t2.y) +
		(t1.z * t2.z) +
		(t1.w * t2.w)
	return
}

func Cross(t1, t2 Tuple) Tuple {
	a := t1.y*t2.z - t1.z*t2.y
	b := t1.z*t2.x - t1.x*t2.z
	c := t1.x*t2.y - t1.y*t2.x
	return Vector(a, b, c)
}

// Represent the projectile after one unit of time has passed
func Tick(e Environment, p Projectile) Projectile {
	position := Add(p.position, p.velocity)
	velocity := Add(Add(p.velocity, e.gravity), e.wind)
	projectile := Projectile{position, velocity}
	return projectile
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
