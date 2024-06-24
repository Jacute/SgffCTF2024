#!/bin/bash

filename="${1%%.*}"

nasm -f elf64 ${filename}".s"
ld ${filename}".o" -o ${filename}

