#!/bin/bash

cd ~/SgffCTF || exit

# down all
for dir in */*/task; do
    if [ -f "$dir/docker-compose.yml" ]; then
        cd "$dir" || exit
        docker-compose down -v
        cd ~/SgffCTF
    fi
done
