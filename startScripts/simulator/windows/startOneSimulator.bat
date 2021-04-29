@echo off
set /p portID="Enter PORT: "
cd Simulator
title %portID%
start "sim 0" cmd /k SimElevatorServer.exe --port %portID%
