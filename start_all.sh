#!/bin/bash

cd ~/SgffCTF || exit

# start pwn
for dir in pwn/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose up -d
        cd ~/SgffCTF
    fi
done

# start web
for dir in web/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose up -d
        cd ~/SgffCTF
    fi
done

# start ppc
for dir in ppc/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose up -d
        cd ~/SgffCTF
    fi
done