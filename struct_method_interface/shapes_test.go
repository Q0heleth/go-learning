package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectange := Rectangle{10.0, 10.0}
	got := Perimeter(rectange)
	want := 40.0
	if got != want {
		t.Errorf("want %.2f, but got %.2f", want, got)
	}
}
func TestArea(t *testing.T) {
	sliceofshape := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{6.0, 7.0}, hasArea: 42.0},
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36},
	}
	for _, tmp := range sliceofshape {
		got := tmp.shape.Area()
		if got != tmp.hasArea {
			t.Errorf("%#v want %.2f, but got %.2f", tmp, tmp.hasArea, got)
		}
	}
	// checkArea:=func (t *testing.T,shape Shape,want float64){
	// 	t.Helper()
	// 	got:=shape.Area()
	// 	if got !=want{
	// 		t.Errorf("want %g, but got %g", want, got)
	// 	}
	// }
	// t.Run("area of rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t,rectangle,72.0)
	// })
	// t.Run("area of circle", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	checkArea(t,circle,314.1592653589793)
	// })
}
