#!/bin/sh

path="$PWD/startScripts/lab"


go build main.go

echo "ID: "
read ID
echo "Enter number of elevators: "
read ELEVATORS

gnome-terminal --window --geometry 50x5+1920+300 --title="Server" -- bash -c "$path/ElevatorServer; bash"
gnome-terminal --window --geometry 200x13+1920+0 --title="Elevator $ID" -- bash -c "./main $ID 4 $ELEVATORS 15657; bash"