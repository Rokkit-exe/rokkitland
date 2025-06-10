#!/bin/bash

echo "Installing Lazy.vim"
if [[ ! -f ~/.config/nvim/lazy-lock.json ]]; then
  git clone https://github.com/LazyVim/starter ~/.config/nvim
  rm -rf ~/.config/nvim/.git
else
  echo "Lazy.vim is already installed"
fi
