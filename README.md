# How to run
To run the program, a server needs to be running first.
You can use either `ElevatorServer` for running in the Real-time Lab, or `SimElevatorServer` for simulation. These are found in respective folders in the [startScripts](./startScripts) folder.  

Then run: `go run main.go ELEVATOR_ID NUM_FLOORS NUM_ELEVATORS ELEVATOR_PORT`  

where `ELEVATOR_ID` is a unique ID starting from 0, `NUM_FLOORS` is the number of floors, `NUM_ELEVATORS` is the number of elevators and `ELEVATOR_PORT` is the port to communicate with the server. Remember to start the server with the same port configuration and use different ports with multiple simulators running on the same time. The `ElevatorServer` port is fixed at 15657.

## Start Scripts
You can also use start scripts found in the [startScripts](./startScripts) folder. These scripts will either start the servers, programs or both.  

# Modules
See picture for how the modules in a single elevator interacts, and the network communication between each elevator.

![Modules](https://folk.ntnu.no/akselna/sanntid.png)
