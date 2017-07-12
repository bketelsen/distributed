#!/bin/bash

docker kill consulsolo
docker rm consulsolo

docker run -d --name consulsolo --net=host -h `hostname` progrium/consul -server -bootstrap -ui-dir /ui
docker ps
