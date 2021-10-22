package main

import (
	"bufio"
	"errors"
	"math"
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
	for _, state := range b.reaction.States() {
		numAtoms := len(state.Atoms())
		distances := make([][]float64, numAtoms)
		for i := range distances {
			distances[i] = make([]float64, numAtoms)
		}

		for i := 0; i < numAtoms; i++ {
			for j := 0; j < numAtoms; j++ {
				if i != j && distances[i][j] == 0 {
					atom1 := state.Atoms()[i]
					atom2 := state.Atoms()[j]

					xDiff := atom1.X() - atom2.X()
					yDiff := atom1.Y() - atom2.Y()
					zDiff := atom1.Z() - atom2.Z()
					dist := math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2) + math.Pow(zDiff, 2))
					distances[i][j] = dist
					distances[j][i] = dist
				}
			}
		}

	}
	return nil
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

		var err error
		if line == "" {
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

	f.WriteString("This is a test output\n")
	return nil
}

func outputState(f *os.File, s State) error {

	return nil
}
