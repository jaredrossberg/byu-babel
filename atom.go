package main

type Atom struct {
	x int
	y int
	z int
}

func createAtom(x, y, z int) Atom {
	return Atom{
		x: x,
		y: y,
		z: z,
	}
}

func (a *Atom) X() int {
	return a.x
}

func (a *Atom) Y() int {
	return a.y
}

func (a *Atom) Z() int {
	return a.z
}
