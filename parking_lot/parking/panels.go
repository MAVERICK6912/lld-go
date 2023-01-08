package parking

import (
	"fmt"
	"time"

	"lld-parking-lot/customerrors"
	"lld-parking-lot/parkingticket"
	"lld-parking-lot/vehicle"

	"github.com/google/uuid"
)

type EntrancePanel struct {
	id                 string
	globalDisplayBoard *ParkingDisplayBoard
	parkingFloors      map[string]*ParkingFloor
}

func NewEntrancePanel(gdb *ParkingDisplayBoard, parkingFloors map[string]*ParkingFloor) *EntrancePanel {
	return &EntrancePanel{
		id:                 uuid.NewString(),
		globalDisplayBoard: gdb,
		parkingFloors:      parkingFloors,
	}
}

func (e EntrancePanel) GetId() string {
	return e.id
}

func (e *EntrancePanel) IssueParkingTicket(v *vehicle.Vehicle) error {
	pf, err := e.checkFloorAvailability(v) //get floor with a parkingSpot available for vehicle
	if err != nil {
		return err
	}
	ps, err := e.checkSpotAvailability(v, pf) // get available parking spot for the vehicle(v) on the parkingfloor(pf)
	if err != nil {
		return err
	}
	pt := parkingticket.ParkingTicket{
		VechicleRegistrationNumber: v.RegistrationNumber,
		IssueTime:                  time.Now(),
		ParkingFloorId:             pf.GetFloorId(),
		ParkingSpotId:              ps.GetSpotId(),
	}
	err = v.AssignTicket(&pt)
	return err
}

func (e EntrancePanel) Error() string {
	return fmt.Sprintf("entrance with id %s already exists", e.id)
}

func (e EntrancePanel) checkFloorAvailability(v *vehicle.Vehicle) (*ParkingFloor, error) {
	for _, pf := range e.parkingFloors {
		if v.IsHandicappedRequired {
			if val, ok := pf.floorDisplayBoard.parkingSpotCount[HandicappedSpot]; ok {
				if val > 0 {
					return pf, nil
				}
				continue
			}
		} else {
			switch v.VehicleType {
			case vehicle.Car:
				if val, ok := pf.floorDisplayBoard.parkingSpotCount[CarSpot]; ok {
					if val > 0 {
						return pf, nil
					}
				}
			case vehicle.Van:
				if val, ok := pf.floorDisplayBoard.parkingSpotCount[LargeSpot]; ok {
					if val > 0 {
						return pf, nil
					}
				}
			case vehicle.Truck:
				if val, ok := pf.floorDisplayBoard.parkingSpotCount[LargeSpot]; ok {
					if val > 0 {
						return pf, nil
					}
				}
			case vehicle.Motorbike:
				if val, ok := pf.floorDisplayBoard.parkingSpotCount[CompactSpot]; ok {
					if val > 0 {
						return pf, nil
					}
				}
			case vehicle.Electric:
				if val, ok := pf.floorDisplayBoard.parkingSpotCount[ElectricSpot]; ok {
					if val > 0 {
						return pf, nil
					}
				}
			}
		}
	}
	return nil, customerrors.ErrorParkingSpotUnavailable
}

func (e EntrancePanel) checkSpotAvailability(v *vehicle.Vehicle, pf *ParkingFloor) (*ParkingSpot, error) {
	// TODO: find a parking spot on the parkingFloor(pf) and assign the vehicle(v) to it.
	if v.IsHandicappedRequired {
		return pf.getFreeParkingSpotForType(HandicappedSpot, v)
	}
	switch v.VehicleType {
	case vehicle.Car:
		return pf.getFreeParkingSpotForType(CarSpot, v)
	case vehicle.Van:
		return pf.getFreeParkingSpotForType(LargeSpot, v)
	case vehicle.Truck:
		return pf.getFreeParkingSpotForType(LargeSpot, v)
	case vehicle.Motorbike:
		return pf.getFreeParkingSpotForType(CompactSpot, v)
	case vehicle.Electric:
		return pf.getFreeParkingSpotForType(ElectricSpot, v)
	default:
		return nil, customerrors.ErrorInvalidVehicleType
	}
}

type ExitPanel struct {
	id            string
	parkingFloors map[string]*ParkingFloor
}

func (e ExitPanel) GetId() string {
	return e.id
}

func NewExitPanel(parkingFloors map[string]*ParkingFloor) *ExitPanel {
	return &ExitPanel{
		id:            uuid.NewString(),
		parkingFloors: parkingFloors,
	}
}

func (e *ExitPanel) AcceptPayment(p *parkingticket.ParkingTicket) error {
	return nil
}

func (e ExitPanel) CalculateAmount(v *vehicle.Vehicle) float64 {
	return v.CaculatePrice()
}

func (e ExitPanel) Error() string {
	return fmt.Sprintf("exit with id %s already exists", e.id)
}
