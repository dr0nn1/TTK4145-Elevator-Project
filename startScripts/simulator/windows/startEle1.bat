@echo off
cd..
cd..
cd..
go build main.go
start "ele 1" cmd /k main.exe 1 4 3 45201
exit