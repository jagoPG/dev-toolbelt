#!/usr/bin/env bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title MD5 a String
# @raycast.mode fullOutput

# @raycast.packageName jagoPG-devkit
# @raycast.argument1 { "type": "text", "placeholder": "the text...", "optional": false }
# @raycast.icon 👾
# @raycast.description Hashes a text with MD5
# @raycast.author Jagoba Pérez-Gómez
# @raycast.authorURL https://jagobapg.eu

echo "$1" | md5
