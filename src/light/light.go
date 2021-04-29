package light

import (
	. "Elevio/elevio"
	. "Src/config"
)

func SetLights(orderCh <-chan OrderList) {

	for {
		select {
		case orderList := <-orderCh:
			updateLights(orderList.Orders)
		}
	}
}

func updateLights(orders [][]Order) {
	for floor, orderTypes := range orders {
		for orderType, thisOrder := range orderTypes {
			lightMode := false
			if thisOrder.State == REQUESTED {
				lightMode = true
			}
			if orderType < BT_Cab || orderType == BT_Cab+ELEV_ID {
				SetButtonLamp(ButtonType(normalizeCabCall(orderType)), floor, lightMode)
			}
		}
	}
}

func normalizeCabCall(btnType int) int {
	if btnType >= BT_Cab {
		btnType = BT_Cab
	}
	return btnType
}
