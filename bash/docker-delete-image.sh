#!/usr/bin/env bash

# Given a Docker image name, deletes all versions of that image

if [[ "$#" != "1" ]]; then
  echo "usage: $0 <IMAGE_NAME>"
  exit 1
fi

imageName=$1
docker images \
        | grep $imageName \
        | tr -s ' ' \
        | cut -d ' ' -f 2 \
        | xargs -I {} docker rmi "$imageName:{}"