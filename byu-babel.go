package main

import (
	"os"
)

type BYUBabelError struct {
	message String
}

func (e *BYUBabelError) Error() string {
	return e.message
}

type BYUBabel struct {
	reaction Reaction
}

func createBabel() BYUBabel {
	return BYUBabel{
		reaction: createReaction(),
	}
}

func (b *BYUBabel) calculate() error {

	return nil
}

func (b *BYUBabel) readFile(inputFile string) error {
	split := strings.SplitN(inputFile, ".", 2)
	if len(split) < 2 {
		return BYUBabelError{"Input file is of unknown filetype"}
	}

	switch split[1] {
	case "xyz":
		return b.readXYZ(inputFile)
	case "mol":
		return b.readMOL(inputFile)
	case "sdf":
		return b.readSDF(inputFile)
	default:
		return BYUBabelError{"Input filetype is not currently accepted"}
	}

	return nil
}

func (b *BYUBabel) readXYZ(inputFile string) error {
	return nil
}

func (b *BYUBabel) readMOL(inputFile string) error {
	return b.readSDF(inputFile)
}

func (b *BYUBabel) readSDF(inputFile string) error {
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

	for _, s := range b.reaction.getStates() {

	}

	f.WriteString("This is a test output\n")
	return nil
}

func outputState(f *os.File, s State) error {

	return nil
}
