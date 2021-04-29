#!/bin/sh

path="$PWD/startScripts/simulator/linux"

gnome-terminal --window --geometry 60x16+1920+0 --title="Sim 0" -- bash -c "$path/SimElevatorServer --port 45200; bash"
gnome-terminal --window --geometry 60x16+1920+400 --title="Sim 1" -- bash -c "$path/SimElevatorServer --port 45201; bash"
gnome-terminal --window --geometry 60x16+1920+800 --title="Sim 2" -- bash -c "$path/SimElevatorServer --port 45202; bash"

go build main.go

$path/startEle0.sh
$path/startEle1.sh
$path/startEle2.sh