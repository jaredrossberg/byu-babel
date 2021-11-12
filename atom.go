package main

import "math"

type Atom struct {
	element string
	x, y, z float64
}

func createAtom(element string, x, y, z float64) Atom {
	return Atom{
		element: element,
		x:       x,
		y:       y,
		z:       z,
	}
}

func (a *Atom) Element() string {
	return a.element
}

func (a *Atom) X() float64 {
	return a.x
}

func (a *Atom) Y() float64 {
	return a.y
}

func (a *Atom) Z() float64 {
	return a.z
}

// UPDATE: This is being removed
// func (a *Atom) getMaxBonds() int {
// 	switch a.element {
// 	case "C":
// 		return 4
// 	case "H":
// 		return 1
// 	default:
// 		return 0
// 	}
// }

func (a *Atom) distanceFrom(a2 Atom) float64 {
	return math.Sqrt(math.Pow(a.x-a2.x, 2) + math.Pow(a.y-a2.y, 2) + math.Pow(a.z-a2.z, 2))
}

func (a *Atom) equals(a2 *Atom) bool {
	return a == a2 ||
		(a.element == a2.element &&
			a.x == a2.x &&
			a.y == a2.y &&
			a.z == a2.z)
}
