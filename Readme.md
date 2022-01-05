# BYU-Babel

This program was created by Brigham Young University to assist with chemistry teaching app [iORA](https://github.com/DanielEss-lab/iORA).

While building iORA, it was found that [OpenBabel](http://openbabel.org/wiki/Main_Page) did not support partial bonds. BYU-Babel aims to replicate the functionality needed from OpenBabel while also supporting partial bonds.

## How to run BYU-Babel

1. [Download](https://www.python.org/downloads/) and install the Python programming language
2. Download or clone this repository
3. Open the repo folder in a terminal or cmd window
3. Run this program using the following example formats: 
    - `python . --input <filename> --output <filename>`
    - `python __main__.py -i <filename> -o <filename>`

Note: (If no output is specified, the output will print to stdout)

## Supported Filetypes
BYU-Babel currently supports the following file formats

Inputs
- xyz
- sdf

Outputs
- sdf
