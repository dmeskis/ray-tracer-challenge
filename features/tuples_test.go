package features

import "testing"

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

func TestPointCreatesPoint(t *testing.T) {
	p := Point(4, -4, 3)

	want := true
	got := IsPoint(p)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestVectorCreatesVector(t *testing.T) {
	v := Vector(4, -4, 3)

	want := true
	got := IsVector(v)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
