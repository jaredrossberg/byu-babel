# Data object representation classes

class Atom:
    def __init__(self, element: str, x: float, y: float, z: float):
        self.element = element
        self.x = x
        self.y = y
        self.z = z

    def distance_from(self, other) -> float :
        return (((self.x - other.x) ** 2) + \
            ((self.y - other.y) ** 2) + \
            ((self.z - other.z) ** 2)) ** 0.5

    def equals(self, other) -> bool:
        return (self == other) or \
            (self.element == other.element and \
            self.x == other.x and \
            self.y == other.y and \
            self.z == other.z)

class Bond:
    def __init__(self, atom1: Atom, atom2: Atom, order: float):
        self.atom1 = atom1
        self.atom2 = atom2
        self.order = order

class State:
    def __init__(self, atoms: list[Atom], bonds: list[Bond] = []):
        self.atoms = atoms
        self.bonds = bonds

    def get_bonds_containing_atom(self, atom: Atom) -> list[Bond]:
        ret = []
        for bond in self.bonds:
            if bond.atom1 == atom or bond.atom2 == atom:
                ret.append(bond)
        return ret
        
    def get_atom_index(self, atom: Atom) -> int:
        for i in range(len(self.atoms)):
            if self.atoms[i].equals(atom):
                return i
        return -1

class Reaction:
    def __init__(self):
        self.states = []

    def addState(self, state: State):
        self.states.append(state)
