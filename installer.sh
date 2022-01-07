#!/bin/bash

#* Creating directories for http-request-creator
mkdir $HOME/.http-request-creator
mkdir $HOME/.http-request-creator/bin

#* Building the app
go build

#* Renaming the app
mv http-request-creator request-creator

#* Moving and copying app and the additional tools to the app directory
mv request-creator $HOME/.http-request-creator/bin
mkdir $HOME/.http-request-creator/tools
cp installer.sh $HOME/.http-request-creator/tools

echo "Now add to the .bashrc or .zshrc command: export PATH=\$PATH:\$HOME/.http-request-creator-bin" #FIXME - check grammar
echo "To use http-request-creator use command: request-creator" #FIXME - check grammar