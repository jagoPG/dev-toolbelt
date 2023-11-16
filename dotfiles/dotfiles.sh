#!/usr/bin/env bash

rm "$HOME/.gitconfig"
rm "$HOME/.gitignore"
rm "$HOME/.editorconfig"
rm "$HOME/.aliases"
rm "$HOME/.vimrc"
rm -r ~/.vim

cp .gitconfig "$HOME/.gitconfig"
cp .gitignore "$HOME/.gitignore"
cp .editorconfig "$HOME/.editorconfig"
cp .aliases "$HOME/.aliases"
cp -r .vim "$HOME/.vim"
cp .vimrc "$HOME/.vimrc"
