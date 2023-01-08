package account

import (
	"lld-parking-lot/parking"
	"lld-parking-lot/parkingticket"
	"lld-parking-lot/vehicle"
)

type attendant Account

type Parkingattendant struct {
	attendant     attendant
	attendantExit *parking.ExitPanel
}

func NewParkingAttendant(exit *parking.ExitPanel, username, pwd string) *Parkingattendant {
	return &Parkingattendant{
		attendant: attendant{
			userName: username,
			password: pwd,
		},
		attendantExit: exit,
	}
}

func (p *Parkingattendant) ProcesssTicket(t *parkingticket.ParkingTicket) error {
	parkinglot := parking.GetParkingLotInstance("alpha")
	return parkinglot.GetExits(p.attendantExit.GetId()).AcceptPayment(t)
}

func (p *Parkingattendant) AssignTicket(e *parking.EntrancePanel, v *vehicle.Vehicle) error {
	parkinglot := parking.GetParkingLotInstance("alpha")
	return parkinglot.GetEntries(e.GetId()).IssueParkingTicket(v)
}
