package main

type State struct {
	atoms []Atom
	bonds []Bond
}

func createState(atoms []Atom, bonds []Bond) State {
	return State{
		atoms: atoms,
		bonds: bonds,
	}
}

func (s *State) getAtoms() []Atom {
	return s.atoms
}

func (s *State) getBonds() []Bond {
	return s.bonds
}
