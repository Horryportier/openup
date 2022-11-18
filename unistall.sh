#! /usr/bin/env bash


opt1="$1"

user=$(whoami)

dir="/home/$user/openup/"

if [[ $opt1 == "-a" ]]; then
        rm -rf $dir
fi
if [[ $opt1 == "-h" ]]; then
        echo -e "\033[0;32mUsage: run {./uninstall.sh} to remove bin. use -a to remove all data\033[0m"
        exit 0
fi

sudo rm "/home/$user/.local/bin/openup"

echo -e "\033[0;33munistalled openup!"
