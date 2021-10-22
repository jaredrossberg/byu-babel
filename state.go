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
