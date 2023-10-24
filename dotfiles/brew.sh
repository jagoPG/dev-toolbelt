#!/usr/bin/env bash
set -e

/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

brew update && brew upgrade

# General tools
brew install coreutils \
  findutils \
  git \
  git-lfs \
  wget \
  grep \
  openssh \
  lynx \
  p7zip \
  rename \
  vbindiff \ # Show differences among files
  zopfli \   # Compression lib
  sqlite \
  tree \
  go

brew install gnu-sed --with-default-names
brew install vim --with-override-system-vi

# Casks
brew install \
  # Workflow
  spectacle \
  raycast \

  # Internet
  firefox \
  chromium \
  spotify \

  # For Development
  iterm2 \
  slack \
  visual-studio-code \
  dbeaver-community \
  macpass \
  docker \
  postman \
  npm \
  neovim \

  # Multimedia
  vlc

brew install zsh-autosuggestions zsh-syntax-highlighting

brew cleanup

echo 'Add "zsh-autosuggestions" and "zsh-syntax-highlighting" to the plugins list at the ".zshrc" configuration file'
