#!/usr/bin/env bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title Count Letters
# @raycast.mode fullOutput

# @raycast.packageName jagoPG-devkit
# @raycast.argument1 { "type": "text", "placeholder": "the text...", "optional": false }
# @raycast.icon 📖
# @raycast.description Count the number of letters
# @raycast.author Jagoba Pérez-Gómez
# @raycast.authorURL https://jagobapg.eu

./../bin/utils count "$1"
