package door

import (
	"Elevio/elevio"
	. "Src/config"
	"time"
)

func Manager(openDoorCh <-chan bool, doorClosedCh chan<- bool) {
	obstruction := false

	obstructionCh := make(chan bool)
	go elevio.PollObstructionSwitch(obstructionCh)

	deadline := make(<-chan time.Time)
	for {
		select {
		case <-openDoorCh:
			deadline = time.After(SERVINGTIME * time.Millisecond)
			elevio.SetDoorOpenLamp(true)

		case <-deadline:
			if !obstruction {
				elevio.SetDoorOpenLamp(false)
				doorClosedCh <- true
			} else {
				deadline = time.After(SERVINGTIME * time.Millisecond)
			}

		case obstruction = <-obstructionCh:
			if !obstruction {
				deadline = time.After(SERVINGTIME / 2 * time.Millisecond)
			}
		}
	}
}
