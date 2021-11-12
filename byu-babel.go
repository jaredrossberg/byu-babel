package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BYUBabel struct {
	reaction Reaction
}

func createBabel() BYUBabel {
	return BYUBabel{
		reaction: createReaction(),
	}
}

func (b *BYUBabel) calculate() error {
	reactionWithBonds := createReaction()
	for _, state := range b.reaction.States() {
		bonds, err := b.calculateSingleState(state)
		if err != nil {
			return err
		}
		state.setBonds(bonds)
		reactionWithBonds.addState(state)
	}
	b.reaction = reactionWithBonds
	return nil
}

func (b *BYUBabel) calculateSingleState(state State) ([]Bond, error) {
	numAtoms := len(state.Atoms())
	distances := make([][]float64, numAtoms)
	bondMatrix := make([][]float64, numAtoms)
	for i := range distances {
		distances[i] = make([]float64, numAtoms)
		bondMatrix[i] = make([]float64, numAtoms)
	}

	// Fill distance matrix with distances
	for i := 0; i < numAtoms; i++ {
		for j := 0; j < numAtoms; j++ {
			if i != j && distances[i][j] == 0 {
				atom1 := state.Atoms()[i]
				atom2 := state.Atoms()[j]

				dist := atom1.distanceFrom(atom2)
				distances[i][j] = dist
				distances[j][i] = dist
			}
		}
	}

	// Fill bond matrix with bonds
	for i := 0; i < numAtoms; i++ {
		for j := 0; j < numAtoms; j++ {
			if i == j || bondMatrix[i][j] != 0 {
				continue
			}
			atom1 := state.Atoms()[i]
			atom2 := state.Atoms()[j]
			dist := distances[i][j]

			if atom1.Element() == "C" && atom2.Element() == "C" {
				if 0 < dist && dist <= 1.24 {
					bondMatrix[i][j] = 3
				} else if 1.24 < dist && dist <= 1.28 {
					bondMatrix[i][j] = 2.5
				} else if 1.28 < dist && dist <= 1.36 {
					bondMatrix[i][j] = 2
				} else if 1.36 < dist && dist <= 1.44 {
					bondMatrix[i][j] = 1.5
				} else if 1.44 < dist && dist <= 1.61 {
					bondMatrix[i][j] = 1
				} else if 1.61 < dist && dist <= 1.89 {
					bondMatrix[i][j] = 0.5
				}
			} else if atom1.Element() == "C" && atom2.Element() == "H" {
				if 0 < dist && dist <= 1.14 {
					bondMatrix[i][j] = 1
				} else if 1.14 < dist && dist <= 1.37 {
					bondMatrix[i][j] = 0.5
				}
			} else if atom1.Element() == "H" && atom2.Element() == "H" {
				if 0 < dist && dist <= 1.63 {
					bondMatrix[i][j] = 1
				} else if 1.63 < dist && dist < 1.00 {
					bondMatrix[i][j] = 0.5
				}
			}
			bondMatrix[j][i] = bondMatrix[i][j]
		}
	}

	// If atom has invalid number of bonds, remove bonds with farthest distance between atoms until valid number of bonds
	// UPDATE: This is being removed
	// for i := 0; i < numAtoms; i++ {
	// 	bondsCount := 0.0
	// 	for j := 0; j < numAtoms; j++ {
	// 		bondsCount += bondMatrix[i][j]
	// 	}
	// 	for bondsCount > float64(state.Atoms()[i].getMaxBonds()) {
	// 		farthestDistance := -1
	// 		for j := 0; j < numAtoms; j++ {
	// 			if bondMatrix[i][j] > 0 && (farthestDistance < 0 || distances[i][j] > distances[i][farthestDistance]) {
	// 				farthestDistance = j
	// 			}
	// 		}
	// 		bondsCount -= bondMatrix[i][farthestDistance]
	// 		bondMatrix[i][farthestDistance] = 0
	// 		bondMatrix[farthestDistance][i] = 0
	// 	}
	// }

	// Add bonds to state
	bonds := make([]Bond, 0)
	for i := 0; i < numAtoms; i++ {
		for j := i + 1; j < numAtoms; j++ {
			if bondMatrix[i][j] == 0 {
				continue
			}
			atom1 := state.Atoms()[i]
			atom2 := state.Atoms()[j]
			order := int(bondMatrix[i][j])
			partial := (bondMatrix[i][j] - float64(order)) > 0
			bonds = append(bonds, createBond(&atom1, &atom2, order, partial))
		}
	}
	return bonds, nil
}

func (b *BYUBabel) readFile(inputFile string) error {
	split := strings.SplitN(inputFile, ".", 2)
	if len(split) < 2 {
		return errors.New("input file is of unknown filetype")
	}

	f, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	switch split[1] {
	case "xyz":
		return b.readXYZ(f)
	case "mol":
		return b.readMOL(f)
	case "sdf":
		return b.readSDF(f)
	default:
		return errors.New("input filetype is not currently accepted")
	}
}

func (b *BYUBabel) readXYZ(f *os.File) error {
	numElements := -1
	count := 0

	atoms := make([]Atom, 0)
	r := createReaction()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "\t")
		if len(words) == 1 {
			words = strings.Split(line, " ")
		}

		var err error
		if line == "" {
			continue
		} else if len(words) > 4 && count == 0 && numElements > 0 {
			continue
		} else if numElements < 0 || (len(words) == 1 && count == 0) {
			var num int
			if num, err = strconv.Atoi(words[0]); err != nil {
				return errors.New("malformed file")
			}
			numElements = num
			continue
		} else if count < numElements {
			var x, y, z float64
			if len(words) < 4 {
				return errors.New("malformed file")
			}
			if x, err = strconv.ParseFloat(words[1], 64); err != nil {
				return errors.New("malformed file")
			}
			if y, err = strconv.ParseFloat(words[2], 64); err != nil {
				return errors.New("malformed file")
			}
			if z, err = strconv.ParseFloat(words[3], 64); err != nil {
				return errors.New("malformed file")
			}
			atoms = append(atoms, createAtom(words[0], x, y, z))
			count++
		}

		if count == numElements {
			r.addState(createStateWithoutBonds(atoms))
			count = 0
			atoms = make([]Atom, 0)
		}

	}
	b.reaction = r
	return nil
}

func (b *BYUBabel) readMOL(f *os.File) error {
	return b.readSDF(f)
}

func (b *BYUBabel) readSDF(f *os.File) error {
	return nil
}

func (b *BYUBabel) outputReaction(outputFile string) error {
	split := strings.SplitN(outputFile, ".", 2)
	if len(split) < 2 || split[1] != "sdf" {
		return errors.New("output filetype is not currently accepted")
	}

	var f *os.File
	if outputFile != "" {
		var err error
		if f, err = os.Create(outputFile); err != nil {
			return err
		}
		defer f.Close()
	} else {
		f = os.Stdout
	}

	for _, s := range b.reaction.States() {
		outputState(f, s)
	}
	return nil
}

func outputState(f *os.File, state State) error {
	f.WriteString("\n")
	f.WriteString("\n")
	f.WriteString("\n")
	f.WriteString(fmt.Sprintln(len(state.Atoms()), len(state.bonds)))
	for _, atom := range state.Atoms() {
		f.WriteString(fmt.Sprintln(atom.X(), atom.Y(), atom.Z(), atom.Element()))
	}
	for _, bond := range state.Bonds() {
		partial := 0
		if bond.partial {
			partial = 1
		}
		f.WriteString(fmt.Sprintln(state.getAtomIndex(bond.atom1)+1, state.getAtomIndex(bond.atom2)+1, bond.order, partial))
		// order := float64(bond.order)
		// if bond.partial {
		// 	order += 0.5
		// }
		// f.WriteString(fmt.Sprintln(state.getAtomIndex(bond.atom1)+1, state.getAtomIndex(bond.atom2)+1, order))
	}
	f.WriteString("M  END\n")
	f.WriteString("$$$$\n")
	return nil
}
