package main

type Bond struct {
	atom1, atom2 *Atom
	order        int
	partial      bool
}

func createBond(atom1, atom2 *Atom, order int, partial bool) Bond {
	return Bond{
		atom1:   atom1,
		atom2:   atom2,
		order:   order,
		partial: partial,
	}
}

func (b *Bond) getAtoms() (*Atom, *Atom) {
	return b.atom1, b.atom2
}

func (b *Bond) getOrder() int {
	return b.order
}

func (b *Bond) hasPartial() bool {
	return b.partial
}
