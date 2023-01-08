package account

import "lld-parking-lot/parking"

type Admin Account

var parkingLot *parking.Parkinglot

func NewAdmin(username, pwd string) *Admin {
	parkingLot = parking.GetParkingLotInstance("alpha")
	return &Admin{
		userName: username,
		password: pwd,
	}
}

func (a *Admin) AddParkingFloor(pf *parking.ParkingFloor) error {
	return parkingLot.AddFloor(pf)
}

func (a *Admin) AddParkingSpot(floorId string, ps *parking.ParkingSpot) error {
	return parkingLot.AddParkingSpot(floorId, ps)
}

func (a *Admin) AddParkingDisplayBoard() error {
	return nil
}

func (a *Admin) AddEntrancePanel(e *parking.EntrancePanel) error {
	return parkingLot.AddEntrance(e)
}

func (a *Admin) AddExitPanel(e *parking.ExitPanel) error {
	return parkingLot.AddExit(e)
}

func (a Admin) GetGlobalDisplayBoard() *parking.ParkingDisplayBoard {
	return parkingLot.GetGlobalDisplayBoard()
}

func (a Admin) GetParkingFloorMap() map[string]*parking.ParkingFloor {
	return a.GetParkingFloorMap()
}
