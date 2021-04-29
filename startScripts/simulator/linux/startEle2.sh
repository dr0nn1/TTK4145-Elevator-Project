#!/bin/sh

go build main.go

gnome-terminal --window --geometry 200x16+0+800 --title="Elevator 2" -- bash -c "./main 2 4 3 45202; exec bash"