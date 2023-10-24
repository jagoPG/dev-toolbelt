#!/usr/bin/env bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title Hex2RGB
# @raycast.mode fullOutput

# @raycast.packageName jagoPG-devkit
# @raycast.argument1 { "type": "text", "placeholder": "HEX", "optional": false }
# @raycast.icon 🎨
# @raycast.description Convert a HEX color to RGB
# @raycast.author Jagoba Pérez-Gómez
# @raycast.authorURL https://jagobapg.eu

./../bin/utils hex2rgb "#$1"
