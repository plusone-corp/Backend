#!/bin/bash

if [[ "$1" == "start" ]]; then
    # Build
    echo "--- BUILDING ---"

    docker build ./ -t plusone-image

    echo "--- FISNIHED BUILDING ---"

    # Run
    echo "--- EXECUTING IMAGES ---"

    docker run -d --name plusone -p 80:80/tcp plusone-image

    echo "--- FINSIHED EXECUTION ---"

    # Confirm
    echo "--- DOCKER PROCESS ARE RUNNING ---"
    echo "visit http://10.8.2.75:8080 for the results"
elif [[ "$1" == "restart" ]]; then
    # Build
    echo "--- DELETING OLD PROCESSES ---"
    docker rm -f plusone

    echo "--- DELETING OLD IMAGES ---"
    docker rmi -f plusone-image

    echo "--- BUILDING ---"

    docker build ./ -t plusone-image

    echo "--- FISNIHED BUILDING ---"

    # Run
    echo "--- EXECUTING IMAGES ---"

    docker run -d --name plusone -p 80:80/tcp plusone-image

    echo "--- FINSIHED EXECUTION ---"

    # Confirm
    echo "--- DOCKER PROCESS ARE RUNNING ---"
    echo "visit http://194.13.83.57:80 for the results"
elif [[ "$1" == "delete" ]]; then
    echo "--- DELETING OLD PROCESSES ---"
    docker rm -f plusone

    echo "--- DELETING OLD IMAGES ---"
    docker rmi -f plusone-image
else
    echo "INVALID OPTIONS"
fi