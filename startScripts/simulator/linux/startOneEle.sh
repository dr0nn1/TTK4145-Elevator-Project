#! /bin/sh

echo "ID: "
read ID
echo "Enter floors: "
read FLOORS
echo "Enter number of elevators: "
read ELEVATORS
echo "Enter PORT: "
read PORT

go build main.go

gnome-terminal --window --title="Elevator $ID" -- bash -c "./main $ID $FLOORS $ELEVATORS $PORT; exec bash"




