package elevator

import (
	"Elevio/elevio"
	. "Src/config"
	"time"
)

func createElevator() Elevator {
	e := Elevator{Direction: DOWN, Floor: -1, State: MOVING, Id: ELEV_ID}
	e.Requests = make([][]bool, NUMFLOORS)
	for i := range e.Requests {
		e.Requests[i] = make([]bool, 3)
	}
	return e
}

func Controller(
	requestCh chan Request,
	doorClosedCh <-chan bool,
	openDoorCh chan<- bool,
	peerCh chan<- Elevator) {

	atFloorCh := make(chan int)
	go elevio.PollFloorSensor(atFloorCh)

	e := createElevator()
	elevio.SetMotorDirection(elevio.MD_Down)

	updateTicker := time.NewTicker(UPDATE_PERIOD)
	for {
		select {
		case request := <-requestCh:
			e.Requests[request.Floor][request.ReqType] = true
			if e.State != IDLE {
				continue
			}
			if request.Floor == e.Floor {
				openDoorCh <- true
				e.State = SERVING_FLOOR
			} else {
				e = ChooseDirection(e)
				moveElevatorInDirection(e)
			}

		case floor := <-atFloorCh:
			elevio.SetFloorIndicator(floor)
			e.Floor = floor
			if ShouldStop(e) {
				elevio.SetMotorDirection(elevio.MD_Stop)
				openDoorCh <- true
				e.State = SERVING_FLOOR
			}

		case <-doorClosedCh:
			e = ClearRequestsAtCurrentFloor(e, requestCh, false)
			e = ChooseDirection(e)
			e = ClearRequestsAtCurrentFloor(e, requestCh, false)
			moveElevatorInDirection(e)

		case <-updateTicker.C:
			peerCh <- e
		}
	}
}

func ShouldStop(e Elevator) bool {
	floor := e.Floor
	if floor == 0 || floor == NUMFLOORS-1 || e.Requests[floor][elevio.BT_Cab] {
		return true
	}
	switch e.Direction {
	case DOWN:
		if e.Requests[floor][elevio.BT_HallDown] || !requestsBelow(e) {
			return true
		}
	case UP:
		if e.Requests[floor][elevio.BT_HallUp] || !requestsAbove(e) {
			return true
		}
	}
	return false
}

func ClearRequestsAtCurrentFloor(e Elevator, requestCh chan<- Request, simulation bool) Elevator {
	for btn := elevio.BT_HallUp; btn <= elevio.BT_Cab; btn++ {
		if !e.Requests[e.Floor][btn] {
			continue
		}
		if int(e.Direction) == int(btn) || btn == elevio.BT_Cab || e.State == IDLE {

			e.Requests[e.Floor][btn] = false
			if !simulation {
				requestCh <- Request{Floor: e.Floor, ReqType: btn}
			}
		}
	}
	return e
}

func moveElevatorInDirection(e Elevator) {
	if e.State == MOVING {
		switch e.Direction {
		case DOWN:
			elevio.SetMotorDirection(elevio.MD_Down)
		case UP:
			elevio.SetMotorDirection(elevio.MD_Up)
		}
	}
}

func ChooseDirection(e Elevator) Elevator {
	e.State = MOVING

	switch e.Direction {
	case DOWN:
		if requestsBelow(e) {
			e.Direction = DOWN
		} else if requestsAbove(e) {
			e.Direction = UP
		} else {
			e.State = IDLE
		}
	case UP:
		if requestsAbove(e) {
			e.Direction = UP
		} else if requestsBelow(e) {
			e.Direction = DOWN
		} else {
			e.State = IDLE
		}
	}
	return e
}

func requestsBelow(e Elevator) bool {
	for i := 0; i < e.Floor; i++ {
		if e.Requests[i][elevio.BT_HallUp] || e.Requests[i][elevio.BT_HallDown] || e.Requests[i][elevio.BT_Cab] {
			return true
		}
	}
	return false
}

func requestsAbove(e Elevator) bool {
	for i := e.Floor + 1; i <= NUMFLOORS-1; i++ {
		if e.Requests[i][elevio.BT_HallUp] || e.Requests[i][elevio.BT_HallDown] || e.Requests[i][elevio.BT_Cab] {
			return true
		}
	}
	return false
}
