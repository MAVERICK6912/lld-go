package parking

import (
	"fmt"
	"lld-parking-lot/customerrors"
	"lld-parking-lot/vehicle"
	"sync"

	"github.com/google/uuid"
)

type ParkingSpot struct {
	id           string
	isFree       bool
	vehicle      *vehicle.Vehicle
	spotType     ParkingSpotType
	parkingFloor *ParkingFloor
	rl           sync.Mutex
}

func NewParkingSpot(pf *ParkingFloor, st ParkingSpotType) *ParkingSpot {
	return &ParkingSpot{
		id:           uuid.NewString(),
		isFree:       true,
		spotType:     st,
		parkingFloor: pf,
	}
}

type ParkingSpotType int

const (
	HandicappedSpot ParkingSpotType = iota
	CompactSpot
	LargeSpot
	ElectricSpot
	CarSpot
)

func (ps *ParkingSpot) AssignSpotToVehicle(v *vehicle.Vehicle) error {
	if ps.vehicle != nil {
		return customerrors.ErrorVehicleAlreadyAssignedToSpot
	}
	ps.isFree = false
	ps.vehicle = v
	ps.parkingFloor.floorDisplayBoard.ChangeCount(&ps.spotType, -1)
	ps.parkingFloor.globalDisplayBoard.ChangeCount(&ps.spotType, -1)
	return nil
}

func (ps *ParkingSpot) FreeSpot() error {
	if ps.vehicle == nil {
		return customerrors.ErrorVehicleNotAssignedToSpot
	}
	ps.isFree = true
	ps.vehicle = nil
	ps.parkingFloor.floorDisplayBoard.ChangeCount(&ps.spotType, 1)
	ps.parkingFloor.globalDisplayBoard.ChangeCount(&ps.spotType, 1)
	return nil
}

func (ps ParkingSpot) Error() string {
	return fmt.Sprintf("parking spot with id %s already exists", ps.id)
}

func (p ParkingSpotType) String() string {
	switch p {
	case HandicappedSpot:
		return "Handicapped Spot"
	case CompactSpot:
		return "Compact Spot"
	case LargeSpot:
		return "Large Spot"
	case ElectricSpot:
		return "Electric Spot"
	case CarSpot:
		return "Car Spot"
	}
	return "unknown"
}

func (ps ParkingSpot) getSpotTypeForVehicle(v vehicle.Vehicle) ParkingSpotType {
	if v.IsHandicappedRequired {
		return HandicappedSpot
	}
	switch v.VehicleType {
	case vehicle.Car:
		return CarSpot
	case vehicle.Van:
		return LargeSpot
	case vehicle.Truck:
		return LargeSpot
	case vehicle.Motorbike:
		return CompactSpot
	case vehicle.Electric:
		return ElectricSpot
	default:
		return CarSpot
	}
}

func (ps ParkingSpot) GetSpotId() string {
	return ps.id
}
