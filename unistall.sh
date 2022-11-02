#! /usr/bin/env bash

user=$(whoami)

dir="/home/$user/openup/"

rm -rf $dir

sudo rm "/home/$user/.local/bin/openup"

echo -e "\033[0;33munistalled openup!"
