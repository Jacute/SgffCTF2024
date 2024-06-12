#!/bin/sh
string=$1
nasm -f elf64 "$string"
gcc -m64 -nostdlib -static "${string%%.*}.o" -o "${string%%.*}"
rm "${string%%.*}.o"
