package parkingticket

import "time"

type ParkingTicket struct {
	VechicleRegistrationNumber string
	IssueTime                  time.Time
	ParkingFloorId             string
	ParkingSpotId              string
}
