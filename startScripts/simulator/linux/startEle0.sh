#!/bin/sh

go build main.go

gnome-terminal --window --geometry 200x16+0+0 --title="Elevator 0" -- bash -c "./main 0 4 3 45200; exec bash"
