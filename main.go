/*
Created by Jared Rossberg on October 20, 2021

This program takes a file and determines bond order (including partial bonds)
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	inputFile := ""
	outputFile := ""

	if len(args) < 1 {
		fmt.Fprint(os.Stderr, "Input file must be provided\n")
		os.Exit(1)
	}
	for i := 0; i < len(args); i++ {
		if args[i] == "-i" || args[i] == "--input" {
			if inputFile != "" {
				fmt.Fprint(os.Stderr, "Input flag cannot be provided more than once\n")
				os.Exit(1)
			} else if i+1 >= len(args) {
				fmt.Fprint(os.Stderr, "Input file must be provided\n")
				os.Exit(1)
			} else {
				inputFile = args[i+1]
				i++
			}
		} else if args[i] == "-o" || args[i] == "--output" {
			if outputFile != "" {
				fmt.Fprint(os.Stderr, "Output flag cannot be provided more than once\n")
				os.Exit(1)
			} else if i+1 >= len(args) {
				fmt.Fprint(os.Stderr, "Output file not provided. Defaulting to stdout\n")
			} else {
				outputFile = args[i+1]
				i++
			}
		}
	}
	if inputFile == "" {
		fmt.Fprint(os.Stderr, "Input file must be provided\n")
		os.Exit(1)
	}

	babel := createBabel()
	if err := babel.readFile(inputFile); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err := babel.calculate(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err := babel.outputReaction(outputFile); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
