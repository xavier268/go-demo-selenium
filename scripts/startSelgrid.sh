#!/bin/bash

echo "Starting a docker stack with the browser"
echo "Type : **docker stack rm selgrid** to remove the stack"

docker --version
docker service ls

docker stack deploy --compose-file selgrid.yaml selgrid
