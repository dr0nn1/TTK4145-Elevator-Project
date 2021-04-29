module Project

require Network-go v0.0.0
replace Network-go => ./dependencies/network

require Elevio v0.0.0
replace Elevio => ./dependencies/elevio

require Src v0.0.0
replace Src => ./src

go 1.16
