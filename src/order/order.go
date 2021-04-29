package order

import (
	"Elevio/elevio"
	. "Src/config"
	"time"
)

var orderList OrderList

func initOrderList() {
	orderList.Owner = ELEV_ID
	orderList.Orders = make([][]Order, NUMFLOORS)
	for i := range orderList.Orders {
		orderList.Orders[i] = make([]Order, elevio.BT_Cab+NUMELEVATORS)
		for j := range orderList.Orders[i] {
			orderList.Orders[i][j].Assignee = -1
		}
	}
}

func Handler(
	requestCh chan Request,
	updateLightsCh chan<- OrderList,
	networkCh chan OrderList,
	costReqCh chan CostRequest,
) {
	buttonCh := make(chan elevio.ButtonEvent)
	go elevio.PollButtons(buttonCh)

	initOrderList()
	updateTicker := time.NewTicker(UPDATE_PERIOD)

	for {
		select {
		case button := <-buttonCh:
			floor := button.Floor
			reqType := unNormalizeCabCall(int(button.Button))
			if orderList.Orders[floor][reqType].State == NONE {
				processBtnPress(floor, reqType, costReqCh)
			}

		case costReq := <-costReqCh:
			floor, reqType, assignee := costReq.Request.Floor, costReq.Request.ReqType, costReq.Assignee
			if orderList.Orders[floor][reqType].State == NONE {
				orderList.Orders[floor][reqType] = Order{Deadline: time.Now().Add(DEADLINE), State: ASSIGNED, Assignee: assignee}
			}

		case remoteOrderList := <-networkCh:
			updateOrdersFromNetwork(remoteOrderList)

		case completedRequest := <-requestCh:
			completeOrder(completedRequest)

		case <-updateTicker.C:
			updateLightsCh <- orderList
			networkCh <- orderList
			processLocalOrderList(requestCh)
		}
	}
}

func processBtnPress(floor int, reqType int, costReqCh chan<- CostRequest) {
	if reqType < elevio.BT_Cab {
		request := Request{Floor: floor, ReqType: elevio.ButtonType(reqType)}
		costReq := CostRequest{Request: request, Assignee: -1}
		costReqCh <- costReq
	} else {
		orderList.Orders[floor][reqType] = Order{Deadline: time.Now().Add(DEADLINE), State: ASSIGNED, Assignee: ELEV_ID}
	}
}

func updateOrdersFromNetwork(remoteOrderList OrderList) {
	listOwner := remoteOrderList.Owner
	if listOwner == ELEV_ID {
		return
	}

	for floor, orderTypes := range remoteOrderList.Orders {
		for orderType, remoteOrder := range orderTypes {
			localOrder := orderList.Orders[floor][orderType]

			switch localOrder.State {
			case NONE:
				switch remoteOrder.State {
				case ASSIGNED:
					localOrder = remoteOrder
				case REQUESTED:
					if listOwner == remoteOrder.Assignee {
						localOrder = remoteOrder
					}
				}

			case ASSIGNED:
				switch remoteOrder.State {
				case REQUESTED:
					if listOwner == localOrder.Assignee {
						localOrder.State = remoteOrder.State
					}
				}

			case REQUESTED:
				switch remoteOrder.State {
				case NONE:
					if listOwner == localOrder.Assignee && listOwner == remoteOrder.Assignee {
						localOrder = remoteOrder
					}
				}
			}

			orderList.Orders[floor][orderType] = localOrder
		}
	}
}

func completeOrder(request Request) {
	floor := request.Floor
	reqType := unNormalizeCabCall(int(request.ReqType))
	orderList.Orders[floor][reqType].Assignee = ELEV_ID
	orderList.Orders[floor][reqType].State = NONE
}

func processLocalOrderList(requestCh chan<- Request) {
	for floor, orderTypes := range orderList.Orders {
		for orderType, localOrder := range orderTypes {
			switch localOrder.State {
			case ASSIGNED:
				if localOrder.Assignee == ELEV_ID {
					requestCh <- Request{Floor: floor, ReqType: elevio.ButtonType(normalizeCabCall(orderType))}
					localOrder.State = REQUESTED
				} else if time.Now().After(localOrder.Deadline) && (normalizeCabCall(orderType) != elevio.BT_Cab) {
					requestCh <- Request{Floor: floor, ReqType: elevio.ButtonType(normalizeCabCall(orderType))}
					localOrder.Assignee = ELEV_ID
					localOrder.State = REQUESTED
					localOrder.Deadline = localOrder.Deadline.Add(DEADLINE)
				}

			case REQUESTED:
				if time.Now().Before(localOrder.Deadline) {
					continue
				}

				if normalizeCabCall(orderType) != elevio.BT_Cab {
					if localOrder.Assignee != ELEV_ID && localOrder.Assignee >= 0 {
						localOrder.Assignee = ELEV_ID
						localOrder.State = REQUESTED
						localOrder.Deadline = localOrder.Deadline.Add(DEADLINE)
						requestCh <- Request{Floor: floor, ReqType: elevio.ButtonType(normalizeCabCall(orderType))}
					} else {
						localOrder.Assignee = -1
					}
				}

				if normalizeCabCall(orderType) == elevio.BT_Cab && localOrder.Assignee != ELEV_ID {
					localOrder.State = ASSIGNED
					localOrder.Deadline = localOrder.Deadline.Add(DEADLINE)
				}
			}

			orderList.Orders[floor][orderType] = localOrder
		}
	}
}

func normalizeCabCall(btnType int) int {
	if btnType >= elevio.BT_Cab {
		btnType = elevio.BT_Cab
	}
	return btnType
}

func unNormalizeCabCall(btnType int) int {
	if btnType == elevio.BT_Cab {
		btnType = elevio.BT_Cab + ELEV_ID
	}
	return btnType
}
