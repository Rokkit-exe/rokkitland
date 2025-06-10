#!/bin/bash

echo "Installing ZSH..."
sudo pacman -S --noconfirm --needed zsh

echo "Setting ZSH as default shell..."
chsh -s $(which zsh)
echo "ZSH installed and set as default shell."

if [[ ! -d ~/.oh-my-zsh ]]; then
  echo "Installing Oh-My-ZSH..."
  sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
  echo "Oh-My-ZSH installed."
else
  echo "Oh-My-ZSH is already installed."
fi

if [[ ! -d ~/.oh-my-zsh/custom/themes/powerlevel10k ]]; then
  echo "Installing Powerlevel10k theme..."
  git clone --depth=1 https://github.com/romkatv/powerlevel10k.git "${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k"
fi

if [[ ! -d ~/.oh-my-zsh/custom/plugins/zsh-autosuggestions ]]; then
  echo "Installing ZSH autosuggestions plugin..."
  git clone https://github.com/zsh-users/zsh-autosuggestions ~/.oh-my-zsh/custom/plugins/zsh-autosuggestions
fi

if [[ ! -d ~/.oh-my-zsh/custom/plugins/zsh-completions ]]; then
  echo "Installing ZSH completions plugin..."
  git clone https://github.com/zsh-users/zsh-completions ~/.oh-my-zsh/custom/plugins/zsh-completions
fi

if [[ ! -d ~/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting ]]; then
  echo "Installing ZSH syntax highlighting plugin..."
  git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
fi

echo "Configuring ZSH..."
sed -i 's/^ZSH_THEME=.*/ZSH_THEME="powerlevel10k/powerlevel10k"/' ~/.zshrc
sed -i 's/^plugins=(.*/plugins=(zsh-autosuggestions zsh-completions zsh-syntax-highlighting z)/' ~/.zshrc
