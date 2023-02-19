package features

import (
	"fmt"
	"math"
	"testing"
)

func TestPoint(t *testing.T) {
	a := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}

	want := true
	got := IsPoint(a)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestVector(t *testing.T) {
	a := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}

	want := true
	got := IsVector(a)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointCreatesPointTuple(t *testing.T) {
	p := Point(4, -4, 3)

	want := true
	got := IsPoint(p)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestVectorCreatesVectorTuple(t *testing.T) {
	v := Vector(4, -4, 3)

	want := true
	got := IsVector(v)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestAddTuple(t *testing.T) {
	p := Point(3, -2, 5)
	v := Vector(-2, 3, 1)

	want := Point(1, 1, 6)
	got := Add(p, v)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSubtractTuple(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	want := Vector(-2, -4, -6)
	got := Subtract(p1, p2)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestNegateTuple(t *testing.T) {
	t1 := Tuple{X: 1.0, Y: -2.0, Z: 3.0, W: -4.0}

	want := Tuple{X: -1.0, Y: 2.0, Z: -3.0, W: 4.0}
	got := Negate(t1)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMultiplyTupleByScalar(t *testing.T) {
	t1 := Tuple{X: 1.0, Y: -2.0, Z: 3.0, W: -4.0}
	scalar := 3.5

	want := Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14}
	got := Multiply(t1, scalar)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMultiplyTupleByFraction(t *testing.T) {
	t1 := Tuple{X: 1.0, Y: -2.0, Z: 3.0, W: -4.0}
	scalar := 0.5

	want := Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}
	got := Multiply(t1, scalar)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDivideTupleByScalar(t *testing.T) {
	t1 := Tuple{X: 1.0, Y: -2.0, Z: 3.0, W: -4.0}
	scalar := 2.0

	want := Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}
	got := Divide(t1, scalar)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMagnitudeWithVectors(t *testing.T) {
	var tests = []struct {
		a, b, c float64
		want    float64
	}{
		{1.0, 0.0, 0.0, 1},
		{0.0, 1.0, 0.0, 1},
		{0.0, 0.0, 1, 1},
		{1.0, 2.0, 3.0, math.Sqrt(14.0)},
		{-1.0, -2.0, -3.0, math.Sqrt(14.0)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v,%v", tt.a, tt.b, tt.c)
		t.Run(testname, func(t *testing.T) {
			tuple := Vector(tt.a, tt.b, tt.c)
			ans := Magnitude(tuple)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestNormalizeVectors(t *testing.T) {
	var tests = []struct {
		a, b, c float64
		want    Tuple
	}{
		{4.0, 0.0, 0.0, Vector(1, 0, 0)},
		{1.0, 2.0, 3.0, Vector(0.26726, 0.53452, 0.80178)},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v,%v", tt.a, tt.b, tt.c)
		t.Run(testname, func(t *testing.T) {
			v := Vector(tt.a, tt.b, tt.c)
			result := Normalize(v)
			if !Equal(result, tt.want) {
				t.Errorf("got %v, want %v", result, tt.want)
			}
		})
	}
}

func TestDotProduct(t *testing.T) {
	t1 := Vector(1.0, 2.0, 3.0)
	t2 := Vector(2.0, 3.0, 4.0)

	want := 20.0
	got := Dot(t1, t2)

	if want != got {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCrossProduct(t *testing.T) {
	t1 := Vector(1.0, 2.0, 3.0)
	t2 := Vector(2.0, 3.0, 4.0)

	want1 := Vector(-1.0, 2.0, -1.0)
	want2 := Vector(1.0, -2.0, 1.0)
	got1 := Cross(t1, t2)
	got2 := Cross(t2, t1)

	if !Equal(want1, got1) {
		t.Errorf("got %v, wanted %v", got1, want1)
	}

	if !Equal(want2, got2) {
		t.Errorf("got %v, wanted %v", got2, want2)
	}
}

// Scenario: Normalizing vector(4, 0, 0) gives (1, 0, 0) Given v ← vector(4, 0, 0)
// Then normalize(v) = vector(1, 0, 0)

// Scenario: Normalizing vector(1, 2, 3)
// Given v ← vector(1, 2, 3)
// Then normalize(v) = approximately vector(0.26726, 0.53452, 0.80178)

// Scenario: The magnitude of a normalized vector Given v ← vector(1, 2, 3)
// When norm ← normalize(v)
// Then magnitude(norm) = 1

// You normalize a tuple by dividing each of its components by its magnitude. In pseudocode, it looks something like this:
// function normalize(v)
// return tuple(v.x / magnitude(v),
// end function
// v.y / magnitude(v),
// v.z / magnitude(v),
// v.w / magnitude(v))

// ch 2
func TestAddColors(t *testing.T) {
	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)

	want := Color(1.6, 0.7, 1.0)
	got := Add(c1, c2)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSubtractColors(t *testing.T) {
	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)

	want := Color(0.2, 0.5, 0.5)
	got := Subtract(c1, c2)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMultiplyColorsByScalar(t *testing.T) {
	c := Color(0.2, 0.3, 0.4)

	want := Color(0.4, 0.6, 0.8)
	got := Multiply(c, 2)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestHadamardProduct(t *testing.T) {
	c1 := Color(1.0, 0.2, 0.4)
	c2 := Color(0.9, 1.0, 0.1)

	want := Color(0.9, 0.2, 0.04)
	got := HadamardProduct(c1, c2)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)

	if c.width != 10 {
		t.Errorf("expect height to be 10, got %d", c.width)
	}

	if c.height != 20 {
		t.Errorf("expect height to be 20, got %d", c.height)
	}

	want := Color(0.0, 0.0, 0.0)
	for _, row := range c.body {
		for _, pixel := range row {
			if !Equal(want, pixel) {
				t.Errorf("got %v, wanted %v", pixel, want)
			}
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := NewCanvas(10, 20)
	defaultColor := Color(0.0, 0.0, 0.0)
	red := Color(1.0, 0.0, 0.0)

	if !Equal(defaultColor, PixelAt(&c, 2, 3)) {
		t.Errorf("failed control case")
	}

	WritePixel(&c, red, 2, 3)

	if !Equal(red, PixelAt(&c, 2, 3)) {
		t.Errorf("pixel not red")
	}
}
