package config

import "time"

const UPDATE_PERIOD = 250 * time.Millisecond
const DEADLINE = 20 * time.Second
const PEER_TOO_OLD = 2 * time.Second
const TRAVEL_TIME_PASSING_FLOOR = 500
const TRAVEL_TIME_BETWEEN_FLOOR = 2000
const SERVINGTIME = 3000

var ELEV_ID int
var NUMFLOORS int
var NUMELEVATORS int

const (
	ORDER_PORT         = 60000
	ELEVATOR_PEER_PORT = 60001
)
