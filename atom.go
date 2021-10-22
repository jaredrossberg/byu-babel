package main

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
