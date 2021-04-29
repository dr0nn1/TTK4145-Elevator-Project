@echo off
cd..
cd..
cd..
go build main.go
start "ele 2" cmd /k main.exe 2 4 3 45202
exit