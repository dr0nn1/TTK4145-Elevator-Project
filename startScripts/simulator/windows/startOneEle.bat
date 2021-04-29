@echo off
set /p ID="Enter ID: "
set /p floors="Enter floors: "
set /p nrEle="Enter number of elevators: "
set /p portID="Enter PORT: "
title %portID%
cd..
cd..
cd..
go build main.go
start "ele 0" cmd /k main.exe %ID% %floors% %nrEle% %portID% 




