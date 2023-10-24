#!/usr/bin/env bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title Base64 Decode
# @raycast.mode fullOutput

# @raycast.packageName jagoPG-devkit
# @raycast.argument1 { "type": "text", "placeholder": "the text...", "optional": false }
# @raycast.icon 👾
# @raycast.description Decodes the text from Base64
# @raycast.author Jagoba Pérez-Gómez
# @raycast.authorURL https://jagobapg.eu

./../bin/utils base64Decode "$1"