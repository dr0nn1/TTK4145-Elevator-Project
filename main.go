package main

import (
	"Elevio/elevio"
	"Network-go/localip"
	"Src/communication"
	. "Src/config"
	"Src/cost"
	"Src/door"
	"Src/elevator"
	"Src/light"
	"Src/order"
	"fmt"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	// Initialize global constants
	ELEV_ID, _ = strconv.Atoi(os.Args[1])
	NUMFLOORS, _ = strconv.Atoi(os.Args[2])
	NUMELEVATORS, _ = strconv.Atoi(os.Args[3])

	driver_addr := "localhost:" + os.Args[4]
	ipAddrStr, err := localip.LocalIP()
	if err != nil {
		fmt.Printf("Exiting with error: %v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Starting elevator %v on IP: %v:%v\n", ELEV_ID, ipAddrStr, os.Args[4])

	doorClosedCh := make(chan bool)
	openDoorCh := make(chan bool)
	requestCh := make(chan Request)
	orderCh := make(chan OrderList)
	lightCh := make(chan OrderList)
	costReqCh := make(chan CostRequest)
	elevatorCh := make(chan Elevator)
	peerCh := make(chan Elevator)

	elevio.Init(driver_addr, NUMFLOORS)

	go door.Manager(openDoorCh, doorClosedCh)
	go order.Handler(requestCh, lightCh, orderCh, costReqCh)
	go light.SetLights(lightCh)
	go elevator.Controller(requestCh, doorClosedCh, openDoorCh, elevatorCh)
	go cost.Assigner(costReqCh, peerCh)
	go communication.NetworkInterface(orderCh, elevatorCh, peerCh)
	go ctrlC()
	select {}
}

func ctrlC() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Print("Exiting...\n")
	elevio.SetMotorDirection(elevio.MD_Stop)
	fmt.Println("Elevator has stopped")
	os.Exit(1)
}
