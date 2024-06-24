#!/bin/bash

cd ~/SgffCTF || exit

# start all
for dir in */*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose up -d
        cd ~/SgffCTF
    fi
done
