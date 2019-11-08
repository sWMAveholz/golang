package main

import ("testing")

func TestPaint(t *testing.T){
	objects := []Painter{
		// short initialization
		Circle{GeoObject{Position{1, 2}, 3}, 40},
		Rectangle{GeoObject{Position{1, 2}, 4}, 10, 10},
		Triangle{GeoObject{Position{1, 2}, 3}, Position{10, 20}, Position{11, 21}, Position{12, 22}},
	}
	for _, v := range objects {
		v.Paint()
	}
}
