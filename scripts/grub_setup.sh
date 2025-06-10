#!/bin/bash

theme_path="/mnt/storage/grub_theme/arch/"
grub_themes_dir="/boot/grub/themes/"
grub_theme_path="$grub_themes_dir/arch/theme.txt"
grub_config="/etc/default/grub"

if [[ -d /mnt/storage/grub_theme/arch ]]; then
  sudo cp -r $theme_path $grub_themes_dir
  sudo sed -i "s|^.*GRUB_DISABLE_RECOVERY=.*$|GRUB_DISABLE_RECOVERY=false|" "$grub_config"
  sudo sed -i "s|^.*GRUB_THEME=.*$|GRUB_THEME=$grub_theme_path|" "$grub_config"
  sudo sed -i "s|^.*GRUB_DISABLE_OS_PROBER=.*$|GRUB_DISABLE_OS_PROBER=false|" "$grub_config"
  sudo grub-mkconfig -o /boot/grub/grub.cfg
else
  echo "Grub theme does not exist in $theme_path"
fi
