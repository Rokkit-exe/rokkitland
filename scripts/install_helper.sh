#!/bin/bash

if [ -x "$(command -v yay)" ]; then
  echo "yay is already installed."
  exit 0
else
  echo "yay is not installed. Proceeding with installation..."
  sudo pacman -S --needed --noconfirm git base-devel
  git clone https://aur.archlinux.org/yay.git
  cd yay
  makepkg -si --noconfirm
  cd ..
  rm -rf yay
fi
