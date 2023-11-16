#!/usr/bin/env bash
set -e

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd -P)

# Install XCode Developer Tools
xcode-select --install

# Link dotfiles
source "$SCRIPT_DIR/dotfiles.sh"

# Install ZSH, HomeBrew and NPM Global Packages
source "$SCRIPT_DIR/ohmyzsh.sh"
source "$SCRIPT_DIR/brew.sh"
source "$SCRIPT_DIR/npm.sh"

# Extras
wget https://download.jetbrains.com/fonts/JetBrainsMono-2.304.zip -O ~/Downloads/JetBrainsMono-2.304.zip
