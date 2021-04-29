cd Simulator
start "sim 0" SimElevatorServer.exe --port 45200
start "sim 1" SimElevatorServer.exe --port 45201
start "sim 2" SimElevatorServer.exe --port 45202
cd ..
start startEle0
start startEle1
start startEle2