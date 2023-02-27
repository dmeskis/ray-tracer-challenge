package features

import (
	"fmt"
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
	Position Tuple // Point
	Velocity Tuple // Vector
}

type Environment struct {
	Gravity Tuple // Vector
	Wind    Tuple // Vector
}

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
	// data [4]interface{}
}

// type Color struct {
// 	Red   float64
// 	Green float64
// 	Blue  float64
// }

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
	t.X = t1.X + t2.X
	t.Y = t1.Y + t2.Y
	t.Z = t1.Z + t2.Z
	t.W = t1.W + t2.W
	// TODO: fix this
	if t.W > 1 {
		panic("O fuk i abded")
	}
	return t
}

func Subtract(t1, t2 Tuple) Tuple {
	t := Tuple{}
	t.X = t1.X - t2.X
	t.Y = t1.Y - t2.Y
	t.Z = t1.Z - t2.Z
	t.W = t1.W - t2.W
	// TODO: fix this
	if t.W < 0 {
		panic("O fuk i subtarcted")

	}
	return t
}

func IsPoint(t Tuple) bool {
	return t.W == 1.0
}

func IsVector(t Tuple) bool {
	return t.W == 0.0
}

func Equal(t1, t2 Tuple) bool {
	return equal(t1.X, t2.X) &&
		equal(t1.Y, t2.Y) &&
		equal(t1.Z, t2.Z) &&
		equal(t1.W, t2.W)
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
	t.X = t1.X * scalar
	t.Y = t1.Y * scalar
	t.Z = t1.Z * scalar
	t.W = t1.W * scalar
	return t
}

func Divide(t1 Tuple, scalar float64) Tuple {
	t := Tuple{}
	t.X = t1.X / scalar
	t.Y = t1.Y / scalar
	t.Z = t1.Z / scalar
	t.W = t1.W / scalar
	return t
}

func Magnitude(t Tuple) float64 {
	radicand := (t.X * t.X) + (t.Y * t.Y) + (t.Z * t.Z) + (t.W * t.W)
	return math.Sqrt(radicand)
}

func Normalize(t1 Tuple) Tuple {
	m := Magnitude(t1)
	t := Tuple{}
	t.X = t1.X / m
	t.Y = t1.Y / m
	t.Z = t1.Z / m
	t.W = t1.W / m
	return t
}

func Dot(t1, t2 Tuple) (scalar float64) {
	scalar = (t1.X * t2.X) +
		(t1.Y * t2.Y) +
		(t1.Z * t2.Z) +
		(t1.W * t2.W)
	return
}

func Cross(t1, t2 Tuple) Tuple {
	a := t1.Y*t2.Z - t1.Z*t2.Y
	b := t1.Z*t2.X - t1.X*t2.Z
	c := t1.X*t2.Y - t1.Y*t2.X
	return Vector(a, b, c)
}

// Represent the projectile after one unit of time has passed
func Tick(e Environment, p Projectile) Projectile {
	position := Add(p.Position, p.Velocity)
	velocity := Add(Add(p.Velocity, e.Gravity), e.Wind)
	projectile := Projectile{position, velocity}
	return projectile
}

// -----------------------------
// Ch 2, rendering the canvas

func Color(r, g, b float64) Tuple {
	return Tuple{
		X: r,
		Y: g,
		Z: b,
	}
}

func HadamardProduct(c1, c2 Tuple) Tuple {
	r := c1.X * c2.X
	g := c1.Y * c2.Y
	b := c1.Z * c2.Z
	return Color(r, g, b)
}

type Canvas struct {
	width  int
	height int
	body   [][]Tuple
}

func (c *Canvas) ToPPM() string {
	// header
	header := ""
	header += "P3\n"
	header += fmt.Sprintf("%d %d\n", c.width, c.height)
	header += "255"

	// body (line should not be more than 70 characters..)
	line := ""
	for _, col := range c.body {
		for _, pixel := range col { // row
			// formula for converting rgb to int
			r := int(math.Floor(Clamp(pixel.X)*255.0 + 0.5))
			g := int(math.Floor(Clamp(pixel.Y)*255.0 + 0.5))
			b := int(math.Floor(Clamp(pixel.Z)*255.0 + 0.5))
			num := fmt.Sprintf("%d %d %d", r, g, b)
			if len(line)+len(num) > 30 {
				line = fmt.Sprintf("%s%s", line, num)
				header += fmt.Sprintf("\n%s", line)
				line = ""
			} else {
				line += fmt.Sprintf("%s ", num)
			}
		}

	}
	header += fmt.Sprintf("%s", line)

	return header
}

func NewCanvas(w, h int) Canvas {
	var body [][]Tuple
	for i := 0; i < h; i++ {
		var row []Tuple
		for k := 0; k < w; k++ {
			row = append(row, Color(0, 0, 0))
		}
		body = append(body, row)
	}
	return Canvas{
		width:  w,
		height: h,
		body:   body,
	}
}

func WritePixel(c *Canvas, color Tuple, x, y int) Canvas {
	c.body[y][x] = color
	return *c
}

func PixelAt(c *Canvas, x, y int) Tuple {
	return c.body[y][x]
}

func Clamp(n float64) float64 {
	if n > 1.0 {
		return 1.0
	} else if n < 0.0 {
		return 0.0

	}
	return n
}
