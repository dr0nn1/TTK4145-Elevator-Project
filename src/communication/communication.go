package communication

import (
	"Network-go/bcast"
	. "Src/config"
)

func NetworkInterface(ordersCh chan OrderList, elevatorCh chan Elevator, peerCh chan Elevator) {
	txOrderCh := make(chan OrderList)
	rxOrderCh := make(chan OrderList)
	txElevatorCh := make(chan Elevator)
	rxElevatorCh := make(chan Elevator)

	go bcast.Transmitter(ORDER_PORT, txOrderCh)
	go bcast.Transmitter(ELEVATOR_PEER_PORT, txElevatorCh)
	go bcast.Receiver(ORDER_PORT, rxOrderCh)
	go bcast.Receiver(ELEVATOR_PEER_PORT, rxElevatorCh)

	for {
		select {
		// To network
		case order := <-ordersCh:
			txOrderCh <- order
		case elevator := <-elevatorCh:
			txElevatorCh <- elevator

		// From network
		case order := <-rxOrderCh:
			ordersCh <- order
		case elevator := <-rxElevatorCh:
			peerCh <- elevator
		}
	}
}
