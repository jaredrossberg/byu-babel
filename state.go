package main

type State struct {
	atoms []Atom
	bonds []Bond
}

func createStateWithBonds(atoms []Atom, bonds []Bond) State {
	return State{
		atoms: atoms,
		bonds: bonds,
	}
}

func createStateWithoutBonds(atoms []Atom) State {
	return State{
		atoms: atoms,
		bonds: make([]Bond, 0),
	}
}

func (s *State) Atoms() []Atom {
	return s.atoms
}

func (s *State) Bonds() []Bond {
	return s.bonds
}

func (s *State) setBonds(bonds []Bond) {
	s.bonds = bonds
}

func (s *State) getBondsContainingAtom(atom Atom) []Bond {
	ret := make([]Bond, 0)
	for i := range s.bonds {
		if s.bonds[i].atom1 == &atom || s.bonds[i].atom2 == &atom {
			ret = append(ret, s.bonds[i])
		}
	}
	return ret
}

func (s *State) getAtomIndex(atom *Atom) int {
	for i := range s.atoms {
		if s.atoms[i].equals(atom) {
			return i
		}
	}
	return -1
}
