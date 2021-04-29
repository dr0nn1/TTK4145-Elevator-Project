#!/bin/sh

path="$PWD/startScripts/simulator/linux"

echo "Elevator port:"
read PORT

gnome-terminal --window -- bash -c "$path/SimElevatorServer --port $PORT; bash"
