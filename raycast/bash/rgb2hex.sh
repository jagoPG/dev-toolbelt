#!/usr/bin/env bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title RGB2Hex
# @raycast.mode fullOutput

# @raycast.packageName jagoPG-devkit
# @raycast.argument1 { "type": "text", "placeholder": "R", "optional": false }
# @raycast.argument2 { "type": "text", "placeholder": "G", "optional": false }
# @raycast.argument3 { "type": "text", "placeholder": "B", "optional": false }
# @raycast.icon 🎨
# @raycast.description Convert a RGB color to HEX
# @raycast.author Jagoba Pérez-Gómez
# @raycast.authorURL https://jagobapg.eu

./../bin/utils rgb2hex $1 $2 $3
