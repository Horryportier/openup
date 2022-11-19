#! /usr/bin/env bash

user="$(whoami)"

path="/home/$user/openup/"

data=$(cat data.json)

config=$(cat config.json)


if [[ -z "$(which go)" ]]; then
        echo -e "\033[0;31mGolang not installed\033[0m"
        exit 1
fi

# update bin file 
go build main.go 
cp main openup

if [[ ! -d "$PATH" ]]; then
        mkdir $path
fi


sudo cp  "$(pwd)/openup" "/home/$user/.local/bin"

cd $path

if [[ ! -e "data.json" ]]; then
        echo $data >> data.json
fi
if [[ ! -e "config.json" ]]; then
        echo $config >> config.json
fi

echo -e "\033[1;32mopenup installed."
