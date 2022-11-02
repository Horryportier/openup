#! /usr/bin/env bash

user="$(whoami)"

path="/home/$user/openup/"

data="{ \"item\":[], \"editor\":\"vim\"}"

if [[ ! -d "$PATH" ]]; then
        $(mkdir $path)
fi


sudo cp  "$(pwd)/openup" "/home/$user/.local/bin"

cd $path

if [[ ! -e "data.json" ]]; then
        echo $data >> data.json
fi
echo -e "\033[1;32mopenup installed."
