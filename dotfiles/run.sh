#!/usr/bin/env bash


# Link dotfiles
source ./dotfiles.sh

# Install XCode Developer Tools
xcode-select --install

# Install ZSH, HomeBrew and NPM Global Packages
source ./ohmyzsh.sh
source ./brew.sh
source ./npm.sh

# Extras
wget https://download.jetbrains.com/fonts/JetBrainsMono-2.304.zip -O ~/Downloads/JetBrainsMono-2.304.zip
