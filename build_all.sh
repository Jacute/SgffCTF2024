#!/bin/bash

cd ~/SgffCTF || exit

# build all
for dir in */*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose build
        cd ~/SgffCTF
    fi
done
