package parking

import (
	"fmt"
	"lld-parking-lot/customerrors"
	"lld-parking-lot/vehicle"
	"sync"

	"github.com/google/uuid"
)

type ParkingFloor struct {
	id                 string
	parkingSpots       map[string]*ParkingSpot
	globalDisplayBoard *ParkingDisplayBoard
	floorDisplayBoard  *ParkingDisplayBoard
	rl                 sync.Mutex
}

func NewParkingFloor(gdb *ParkingDisplayBoard) *ParkingFloor {
	return &ParkingFloor{
		id:                 uuid.NewString(),
		parkingSpots:       make(map[string]*ParkingSpot),
		globalDisplayBoard: gdb,
	}
}

func (pf *ParkingFloor) addParkingSpot(ps *ParkingSpot) error {
	pf.rl.Lock()
	defer pf.rl.Unlock()
	if _, ok := pf.parkingSpots[ps.id]; ok {
		return ps
	}
	pf.parkingSpots[ps.id] = ps
	return nil
}

func (pf *ParkingFloor) RemoveParkingSpot(ps *ParkingSpot) error {
	if _, ok := pf.parkingSpots[ps.id]; ok {
		delete(pf.parkingSpots, ps.id)
		return nil
	}
	return customerrors.ErrorParkingSpotDoesNotExist
}

func (pf ParkingFloor) Error() string {
	return fmt.Sprintf("parking floor with id %s already exists", pf.id)
}

func (pf ParkingFloor) GetFloorId() string {
	return pf.id
}

func (pf ParkingFloor) getFreeParkingSpotForType(st ParkingSpotType, v *vehicle.Vehicle) (*ParkingSpot, error) {
	for _, ps := range pf.parkingSpots {
		ps.rl.Lock()
		if ps.spotType == st {
			if ps.isFree {
				err := ps.AssignSpotToVehicle(v)
				if err != nil {
					return nil, err
				}
				return ps, nil
			}
		}
		ps.rl.Unlock()
	}
	return nil, customerrors.ErrorParkingSpotUnavailable
}
