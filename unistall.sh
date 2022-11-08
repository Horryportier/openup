#! /usr/bin/env bash

# TODO flag to choose deleting data.joson or only bin.

user=$(whoami)

dir="/home/$user/openup/"

rm -rf $dir

sudo rm "/home/$user/.local/bin/openup"

echo -e "\033[0;33munistalled openup!"
