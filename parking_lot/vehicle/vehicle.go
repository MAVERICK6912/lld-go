package vehicle

import (
	"lld-parking-lot/customerrors"
	"lld-parking-lot/parkingticket"
	"time"
)

type Vehicle struct {
	RegistrationNumber    string
	VehicleType           vehicleType
	Ticket                *parkingticket.ParkingTicket
	IsHandicappedRequired bool
}

type vehicleType int

const (
	Car vehicleType = iota
	Truck
	Van
	Motorbike
	Electric
)

func (v *Vehicle) AssignTicket(p *parkingticket.ParkingTicket) error {
	if v.Ticket == nil {
		v.Ticket = p
		return nil
	}
	return customerrors.ErrorTicketAlreadyAssigned
}

func (v Vehicle) CaculatePrice() float64 {
	timeDiff := time.Since(v.Ticket.IssueTime)
	return timeDiff.Hours() * v.VehicleType.getCharge()
}

func (vt vehicleType) String() string {
	switch vt {
	case Car:
		return "Car"
	case Truck:
		return "Truck"
	case Van:
		return "Van"
	case Motorbike:
		return "Motorbike"
	case Electric:
		return "Electric"
	default:
		return "unknown"
	}
}

func (vt vehicleType) Error() string {
	switch vt {
	case Car:
	case Truck:
	case Van:
	case Motorbike:
	case Electric:
	default:
		return "unknown vehicle type"
	}
	return ""
}

func (vt vehicleType) getCharge() float64 {
	switch vt {
	case Car:
		return 15
	case Truck:
		return 40
	case Van:
		return 25
	case Motorbike:
		return 7
	case Electric:
		return 10
	default:
		return 5
	}
}

func (v Vehicle) GetTicket() *parkingticket.ParkingTicket {
	return v.Ticket
}
