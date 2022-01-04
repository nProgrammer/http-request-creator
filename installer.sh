#!/bin/bash

mkdir $HOME/.http-request-creator
mkdir $HOME/.http-request-creator/bin

go build
mv http-request-creator request-creator

mv request-creator $HOME/.http-request-creator/bin

echo "Now add to the .bashrc or .zshrc command: export PATH=\$PATH:\$HOME/.http-request-creator-bin"
echo "To use http-request-creator use command: request-creator"