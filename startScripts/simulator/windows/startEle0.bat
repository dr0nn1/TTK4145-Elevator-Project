@echo off
cd..
cd..
cd..
go build main.go
start "ele 0" cmd /k main.exe 0 4 3 45200
exit
