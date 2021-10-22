package main

type Reaction struct {
	states []State
}

func createReaction() Reaction {
	return Reaction{
		states: make([]State, 0),
	}
}

func (r *Reaction) addState(s State) {
	r.states = append(r.states, s)
}

func (r *Reaction) States() []State {
	return r.states
}
