'''
Created by Jared Rossberg on October 20, 2021

This program takes a file and determines bond order (including partial bonds)
'''

import json
import sys
from dataobjects import Reaction
from dataobjects import State
from dataobjects import Atom
from dataobjects import Bond

class BYUBabel:
    def __init__(self, input_file, config_file, translation):
        self.input_file = input_file
        self.reaction = Reaction()
        self._load_config(config_file)

        file_type = self.input_file.split('.',1)[1]
        if file_type == 'xyz':
            self._parse_xyz(translation)
        elif file_type == 'sdf':
            self._parse_sdf(translation)
        else:
            raise Exception('File type not supported')
        pass

    class Config_Pair:
        def __init__(self, element1: str, element2: str, offset):
            self.element1 = element1
            self.element2 = element2
            self.bond_distances = []
            self.offset = offset
            pass

    def _load_config(self, config_file):
        f = open(config_file)
        self.config = {}
        for pair in json.load(f)['distance-criteria']:
            element1 = pair['element1']
            element2 = pair['element2']
            offset = pair['offset']

            if not element1 in self.config:
                self.config[element1] = {}
            if not element2 in self.config:
                self.config[element2] = {}

            cp = self.Config_Pair(element1, element2, offset)
            for bond in pair['bond-formation']:
                cp.bond_distances.append({
                    'strength': bond['bond-strength'],
                    'min': bond['min-distance'], 
                    'max': bond['max-distance']})
            self.config[element1][element2] = cp
            self.config[element2][element1] = cp
        f.close()

        f = open(config_file)
        self.transition_state = json.load(f)['transition-state']
        f.close()

        f = open(config_file)
        self.window = json.load(f)['window']
        f.close()

        f = open(config_file)
        self.frozen_before = {}
        for frozen_bond in json.load(f)['frozen-before']:
            atom1 = frozen_bond['atom1'] - 1
            atom2 = frozen_bond['atom2'] - 1
            bond_strength = frozen_bond['bond-strength']

            if not atom1 in self.frozen_before:
                self.frozen_before[atom1] = {}
            if not atom2 in self.frozen_before:
                self.frozen_before[atom2] = {}
            
            self.frozen_before[atom1][atom2] = bond_strength
            self.frozen_before[atom2][atom1] = bond_strength
        f.close()

        f = open(config_file)
        self.frozen_forever = {}
        for frozen_bond in json.load(f)['frozen-forever']:
            atom1 = frozen_bond['atom1'] - 1
            atom2 = frozen_bond['atom2'] - 1
            bond_strength = frozen_bond['bond-strength']

            if not atom1 in self.frozen_forever:
                self.frozen_forever[atom1] = {}
            if not atom2 in self.frozen_forever:
                self.frozen_forever[atom2] = {}
            
            self.frozen_forever[atom1][atom2] = bond_strength
            self.frozen_forever[atom2][atom1] = bond_strength
        f.close()

        f = open(config_file)
        self.frozen_after= {}
        for frozen_bond in json.load(f)['frozen-after']:
            atom1 = frozen_bond['atom1'] - 1
            atom2 = frozen_bond['atom2'] - 1
            bond_strength = frozen_bond['bond-strength']

            if not atom1 in self.frozen_after:
                self.frozen_after[atom1] = {}
            if not atom2 in self.frozen_after:
                self.frozen_after[atom2] = {}
            
            self.frozen_after[atom1][atom2] = bond_strength
            self.frozen_after[atom2][atom1] = bond_strength
        f.close()
        pass

    def calculate(self):
        for i, state in enumerate(self.reaction.states):
            if i < self.transition_state - self.window:
                state.bonds = self._calculate_state(state, self.frozen_before, False)
            elif i > self.transition_state + self.window:
                state.bonds = self._calculate_state(state, self.frozen_after, False)
            else:
                state.bonds = self._calculate_state(state, [], True)
        return self

    def _calculate_state(self, state: State, frozen_bonds, is_during_transition) -> list[Bond]:
        num_atoms = len(state.atoms)

        # Fill distance matrix
        distances = [ [0]*num_atoms for _ in range(num_atoms) ]
        for i in range(num_atoms):
            for j in range(num_atoms):
                if i != j and distances[i][j] == 0:
                    dist = state.atoms[i].distance_from(state.atoms[j])
                    distances[i][j] = dist
                    distances[j][i] = dist

        # Fill bond matrix
        bond_matrix = [ [0]*num_atoms for _ in range(num_atoms) ]

        for i in self.frozen_forever:
            for j in self.frozen_forever[i]:
                bond_matrix[i][j] = self.frozen_forever[i][j]
        for i in frozen_bonds:
            for j in frozen_bonds[i]:
                bond_matrix[i][j] = frozen_bonds[i][j]
        if is_during_transition:
            for i in range(num_atoms):
                for j in range(num_atoms):
                    if i == j or bond_matrix[i][j] != 0:
                        continue

                    element1 = state.atoms[i].element
                    element2 = state.atoms[j].element

                    # Handles scenario where config file does not include pairing distances between elements
                    # Currently does not throw an error, and continues silently
                    if not element1 in self.config or not element2 in self.config[element1]:
                        break

                    config_pair = self.config[element1][element2]
                    dist = distances[i][j] - config_pair.offset
                    for x in config_pair.bond_distances:
                        if x['min'] < dist and dist <= x['max']:
                            bond_matrix[i][j] = x['strength']
                            break
                    bond_matrix[j][i] = bond_matrix[i][j]

        # Add bonds to state
        bonds = []
        for i in range(num_atoms):
            for j in range(i+1, num_atoms):
                if bond_matrix[i][j] <= 0:
                    continue
                bonds.append(Bond(state.atoms[i], state.atoms[j], bond_matrix[i][j]))
        return bonds

    def _parse_xyz(self, translation):
        f = open(self.input_file)

        num_elements = -1
        count = 0
        atoms = []

        for line in f.readlines():
            words = line.split('\t')
            if len(words) == 1:
                words = line.split(' ')
            while("" in words) :
                words.remove("")
            
            if line == '' or line == '\n':
                continue
            elif len(words) > 4 and count == 0 and num_elements > 0:
                continue
            elif num_elements < 0 or (len(words) == 1 and count == 0):
                try:
                    num_elements = int(words[0])
                except:
                    self.malformed_file()
                continue
            elif count < num_elements:
                if len(words) < 4:
                    self.malformed_file()
                try:
                    x = float(words[1]) + translation['x']
                    y = float(words[2]) + translation['y']
                    z = float(words[3]) + translation['z']
                    atoms.append(Atom(words[0], x, y, z))
                    count += 1
                except:
                    self.malformed_file()

            if count == num_elements:
                self.reaction.addState(State(atoms))
                count = 0
                atoms = []

        pass

    def _parse_sdf(self, translation):
        f = open(self.input_file)

        ignore_lines_start = 3
        ignore_lines_end = 2
        num_atoms = -1
        num_bonds = -1
        array_atoms = []

        for line in f.readlines():
            words = line.split('\t')
            if len(words) == 1:
                words = line.split(' ')
            while("" in words) :
                words.remove("")

            if ignore_lines_start > 0:
                ignore_lines_start -= 1
            elif num_atoms < 0 and num_bonds < 0:
                if len(words) < 2:
                    self.malformed_file()
                try:
                    num_atoms = int(words[0])
                    num_bonds = int(words[1])
                except:
                    self.malformed_file()
            elif num_atoms > 0:
                if len(words) < 4:
                    self.malformed_file()
                try:
                    element = words[3]
                    x = float(words[0]) + translation['x']
                    y = float(words[1]) + translation['y']
                    z = float(words[2]) + translation['z']
                    array_atoms.append(Atom(element, x, y, z))
                    num_atoms -= 1
                except:
                    self.malformed_file
            elif num_bonds > 0:
                num_bonds -= 1
            elif ignore_lines_end > 0:
                ignore_lines_end -= 1

            if ignore_lines_start == 0 and num_atoms == 0 and num_bonds == 0:
                if ignore_lines_end == 1:
                    self.reaction.addState(State(array_atoms))
                elif ignore_lines_end == 0:
                    ignore_lines_start = 3
                    ignore_lines_end = 2
                    num_atoms = -1
                    num_bonds = -1
                    array_atoms = []
        pass

    def malformed_file(self):
        self._eprint('malformed file')
        sys.exit(1)

    def _eprint(*args, **kwargs):
        print(*args, file=sys.stderr, **kwargs)

    def write_reaction(self, output_file: str = ''):
        original_stdout = sys.stdout
        if output_file == '':
            self._write_reaction_helper()
        else:
            with open(output_file, 'w') as f:
                sys.stdout = f
                self._write_reaction_helper()
                sys.stdout = original_stdout
    
    def _write_reaction_helper(self):
        for state in self.reaction.states:
            self._write_state(state)
        pass

    def _write_state(self, state: State):
        print('\n\n')
        print('{:>3}{:>3}  0  0  0  0  0  0  0  0999 V2000'.format(len(state.atoms), len(state.bonds)))
        for atom in state.atoms:
            print('{:>12f}{:>12f}{:>12f} {:<4}'.format(atom.x, atom.y, atom.z, atom.element))
        for bond in state.bonds:
            print('{:>4}{:>5}{:>5}'.format(state.get_atom_index(bond.atom1)+1, state.get_atom_index(bond.atom2)+1, bond.order))
        print('M  END')
        print('$$$$')
        pass


