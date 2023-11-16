#!/usr/bin/env bash
set -e

/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

brew update && brew upgrade

################ HOMEBREW TOOLS ################
brew install coreutils findutils git git-lfs wget grep openssh lynx p7zip rename vbindiff zopfli sqlite tree go
git lfs install

brew install gnu-sed --with-default-names
brew install vim --with-override-system-vi

# ZSH Tools
brew install zsh-autosuggestions zsh-syntax-highlighting

# Development
brew docker docker-compose gh node yarn


################ HOMEBREW CASKS ################
# Workflow
brew install --cask spectacle raycast

# Internet
brew install --cask firefox google-chrome spotify

# Development
brew install --cask iterm2 slack visual-studio-code dbeaver-community postman npm notion

# Multimedia
brew install --cask vlc

brew cleanup

echo 'Add "zsh-autosuggestions" and "zsh-syntax-highlighting" to the plugins list at the ".zshrc" configuration file'
