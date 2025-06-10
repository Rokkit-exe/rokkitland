#!/bin/bash

sddm_config_dir="/etc/sddm.conf.d"
sddm_config_path="$sddm_config_dir/theme.conf.user"
sddm_theme_dir="/usr/share/sddm/themes"
sddm_theme_path="$sddm_theme_dir/simple-sddm-2"
sddm_background_path="$sddm_theme_path/Backgrounds"
background_path="/mnt/storage/wallpaper/background.jpg"
simple_config="$sddm_theme_path/theme.conf"
background_string="Backgrounds/background.jpg"

yay -S --noconfirm --needed simple-sddm-theme-2-git

if [[ -d $sddm_theme_path ]]; then
  sudo mkdir -p $sddm_config_path
  touch sddm_config_path
  sudo echo "[Theme]" >>sddm_config_path
  sudo echo "Current=simple-sddm-2" >>sddm_config_path
  sudo cp $background_path $sddm_background_path
  sudo sed -i "s|^Background=.*$|Background=$background_string|" "$simple_config"
else
  echo "theme not found"
fi
