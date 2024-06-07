#!/bin/bash

cd ~/SgffCTF || exit

# down pwn
for dir in pwn/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose down -v
        cd ~/SgffCTF
    fi
done

# down web
for dir in web/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose down -v
        cd ~/SgffCTF
    fi
done

# down ppc
for dir in ppc/*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose down -v
        cd ~/SgffCTF
    fi
done