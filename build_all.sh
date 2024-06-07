#!/bin/bash

cd ~/SgffCTF || exit

# build pwn
for dir in pwn/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose build
        cd ~/SgffCTF
    fi
done

# build web
for dir in web/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose build
        cd ~/SgffCTF
    fi
done

# build ppc
for dir in ppc/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose build
        cd ~/SgffCTF
    fi
done