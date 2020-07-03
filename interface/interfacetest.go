package main

import "github.com/sirupsen/logrus"

type Object interface {
	Area() int
	Sum() int
}

type Object2 interface {
	Area() int
}

type Rect struct {
	len   int
	wigth int
}

func (this Rect) Area() int {
	return this.len * this.wigth
}

func (this Rect) Sum() int {
	return this.len + this.wigth
}

type Rect2 struct {
	l int
	w int
}

func (this Rect2) Area() int {
	return this.l * this.w
}

func main() {
	var o Object = Rect{len: 20, wigth: 30}
	a := o.(Object)

	c := Rect2{l: 10, w: 20}
	var o2 Object2 = c

	logrus.Infof("area: %d, sun: %d", o.Area(), o.Sum())
	logrus.Infof("2 area: %d", o2.Area())

	logrus.Infof("area: %d, sun: %d", a.Area(), a.Sum())

	o2 = o
	logrus.Infof("2 area: %d", o2.Area())
}
