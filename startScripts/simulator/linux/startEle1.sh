#!/bin/sh

go build main.go

gnome-terminal --window --geometry 200x16+0+400 --title="Elevator 1" -- bash -c "./main 1 4 3 45201; exec bash"