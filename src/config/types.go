package config

import (
	"Elevio/elevio"
	"time"
)

type CostRequest struct {
	Request  Request
	Assignee int
}

type Peer struct {
	Elevator   Elevator
	LastUpdate time.Time
}

type ElevatorState int

const (
	IDLE ElevatorState = iota
	MOVING
	SERVING_FLOOR
)

type Direction int

const (
	UP Direction = iota
	DOWN
)

type Elevator struct {
	Direction Direction
	Floor     int
	Requests  [][]bool
	State     ElevatorState
	Id        int
}

type Request struct {
	Floor   int
	ReqType elevio.ButtonType
}

type OrderState int

const (
	NONE OrderState = iota
	ASSIGNED
	REQUESTED
)

type Order struct {
	Deadline time.Time
	State    OrderState
	Assignee int
}
type OrderList struct {
	Owner  int
	Orders [][]Order
}
