#!/bin/bash

if [[ "$1" == "start" ]]; then
    # Build
    echo "--- BUILDING ---"

    docker build ./ -t geobot-image

    echo "--- FISNIHED BUILDING ---"

    # Run
    echo "--- EXECUTING IMAGES ---"

    docker run -d --name geobot -p 3000:3000/tcp geobot-image

    echo "--- FINSIHED EXECUTION ---"

    # Confirm
    echo "--- DOCKER PROCESS ARE RUNNING ---"
    echo "visit http://10.8.2.75:8080 for the results"
elif [[ "$1" == "restart" ]]; then
    # Build
    echo "--- DELETING OLD PROCESSES ---"
    docker rm -f geobot

    echo "--- DELETING OLD IMAGES ---"
    docker rmi -f geobot-image

    echo "--- BUILDING ---"

    docker build ./ -t geobot-image

    echo "--- FISNIHED BUILDING ---"

    # Run
    echo "--- EXECUTING IMAGES ---"

    docker run -d --name geobot -p 3000:3000/tcp geobot-image

    echo "--- FINSIHED EXECUTION ---"

    # Confirm
    echo "--- DOCKER PROCESS ARE RUNNING ---"
    echo "visit http://194.13.83.57:3000 for the results"
elif [[ "$1" == "delete" ]]; then
    echo "--- DELETING OLD PROCESSES ---"
    docker rm -f geobot

    echo "--- DELETING OLD IMAGES ---"
    docker rmi -f geobot-image
else
    echo "INVALID OPTIONS"
fi