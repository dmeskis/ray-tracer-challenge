package features

import (
	"fmt"
	"math"
	"testing"
)

func TestPoint(t *testing.T) {
	a := Tuple{x: 4.3, y: -4.2, z: 3.1, w: 1.0}

	want := true
	got := IsPoint(a)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestVector(t *testing.T) {
	a := Tuple{x: 4.3, y: -4.2, z: 3.1, w: 0.0}

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
	t1 := Tuple{x: 1.0, y: -2.0, z: 3.0, w: -4.0}

	want := Tuple{x: -1.0, y: 2.0, z: -3.0, w: 4.0}
	got := Negate(t1)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMultiplyTupleByScalar(t *testing.T) {
	t1 := Tuple{x: 1.0, y: -2.0, z: 3.0, w: -4.0}
	scalar := 3.5

	want := Tuple{x: 3.5, y: -7, z: 10.5, w: -14}
	got := Multiply(t1, scalar)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMultiplyTupleByFraction(t *testing.T) {
	t1 := Tuple{x: 1.0, y: -2.0, z: 3.0, w: -4.0}
	scalar := 0.5

	want := Tuple{x: 0.5, y: -1, z: 1.5, w: -2}
	got := Multiply(t1, scalar)

	if !Equal(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDivideTupleByScalar(t *testing.T) {
	t1 := Tuple{x: 1.0, y: -2.0, z: 3.0, w: -4.0}
	scalar := 2.0

	want := Tuple{x: 0.5, y: -1, z: 1.5, w: -2}
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
