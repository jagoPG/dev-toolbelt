#!/usr/bin/env bash

rm $HOME/.gitconfig
rm $HOME/.gitignore
rm $HOME/.editorconfig
rm $HOME/.aliases
rm $HOME/.vimrc
rm -r ~/.vim

ln -s .gitconfig $HOME/.gitconfig
ln -s .gitignore $HOME/.gitignore
ln -s .editorconfig $HOME/.editorconfig
ln -s .aliases $HOME/.aliases
cp -r .vim $HOME/.vim
cp .vimrc $HOME/.vimrc
