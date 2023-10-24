#!/usr/bin/env bash

# Builds the application and places the binary in the Raycast's bin folder
set -e

go build -o ../raycast/bin
