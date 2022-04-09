#!/bin/bash

# Runs babel program on input folder and prints to the output folder 

for f in inputs/*
do
f=$(basename ${f})
python . -i inputs/$f -o outputs/$f
done

