# BYU-Babel

This program was created by Brigham Young University to assist with chemistry teaching app [iORA](https://github.com/DanielEss-lab/iORA).

While building iORA, it was found that [OpenBabel](http://openbabel.org/wiki/Main_Page) did not support partial bonds. BYU-Babel aims to replicate the functionality needed from OpenBabel while also supporting partial bonds.

## How to run BYU-Babel

1. [Download](https://golang.org) and install the Go programming language
2. Download or clone this repository
3. Compile the project into an executable with the `go build` or `go install` command
4. Run the executable using the following format: `./byu-babel --input <filename> --output <filename>` (If no output is specified, the output will print to stdout)

## Supported Filetypes
BYU-Babel currently supports the following file formats

Inputs
- xyz
- sdf
- mol

Outputs
- sdf
