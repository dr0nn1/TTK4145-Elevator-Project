package cost

import (
	. "Src/config"
	"Src/elevator"
	"time"
)

func Assigner(costReqCh chan CostRequest, peerCh <-chan Elevator) {
	peers := make([]Peer, NUMELEVATORS)

	for {
		select {
		case request := <-costReqCh:
			solution := bestAssignee(request, peers)
			costReqCh <- solution
		case elevator := <-peerCh:
			peer := Peer{Elevator: elevator, LastUpdate: time.Now()}
			peers[elevator.Id] = peer
		}
	}
}

func bestAssignee(costReq CostRequest, peers []Peer) CostRequest {
	leastCost := int(^uint(0) >> 1) //Max int
	assignee := ELEV_ID
	for _, peer := range peers {
		if time.Since(peer.LastUpdate) > PEER_TOO_OLD {
			continue
		}
		peer.Elevator.Requests[costReq.Request.Floor][costReq.Request.ReqType] = true
		cost := costForElevator(peer.Elevator)
		if cost < leastCost {
			leastCost = cost
			assignee = peer.Elevator.Id
		}
	}
	costReq.Assignee = assignee
	return costReq
}

func costForElevator(e Elevator) int {
	duration := 0

	switch e.State {
	case IDLE:
		e = elevator.ChooseDirection(e)
		if e.State == IDLE {
			return duration
		}
	case MOVING:
		duration += TRAVEL_TIME_BETWEEN_FLOOR / 2
		e.Floor += changeFloor(e.Direction)
	case SERVING_FLOOR:
		duration -= SERVINGTIME / 2
	}

	for duration < 1000000 {
		if elevator.ShouldStop(e) {
			e = elevator.ClearRequestsAtCurrentFloor(e, nil, true)
			duration += SERVINGTIME
			e = elevator.ChooseDirection(e)
			e = elevator.ClearRequestsAtCurrentFloor(e, nil, true)
			if e.State == IDLE {
				return duration
			}
		} else {
			duration += TRAVEL_TIME_PASSING_FLOOR
		}
		e.Floor += changeFloor(e.Direction)
		duration += TRAVEL_TIME_BETWEEN_FLOOR
	}
	return duration
}

func changeFloor(dir Direction) int {
	i := 0
	switch dir {
	case DOWN:
		i = -1
	case UP:
		i = 1
	}
	return i
}
