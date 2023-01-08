package parking

import (
	"sync"
)

var parkingLot *Parkinglot = nil
var once sync.Once

type Parkinglot struct {
	name               string
	entrances          map[string]*EntrancePanel
	exits              map[string]*ExitPanel
	parkingFloors      map[string]*ParkingFloor
	globalDisplayBoard *ParkingDisplayBoard
}

func newParkingLotInstance(name string, entrances map[string]*EntrancePanel, exits map[string]*ExitPanel, parkingFloors map[string]*ParkingFloor) *Parkinglot {
	once.Do(func() {
		parkingLot = &Parkinglot{
			name:               name,
			entrances:          entrances,
			exits:              exits,
			parkingFloors:      parkingFloors,
			globalDisplayBoard: NewDisplayBoard(),
		}
	})
	return parkingLot
}

func GetParkingLotInstance(name string) *Parkinglot {
	if parkingLot == nil {
		return newParkingLotInstance(name, make(map[string]*EntrancePanel), make(map[string]*ExitPanel), make(map[string]*ParkingFloor))
	}
	return parkingLot
}

func (p *Parkinglot) AddEntrance(e *EntrancePanel) error {
	e.globalDisplayBoard = p.globalDisplayBoard
	if _, ok := p.entrances[e.id]; ok {
		return e
	}
	p.entrances[e.id] = e
	return nil
}

func (p *Parkinglot) AddExit(e *ExitPanel) error {
	if _, ok := p.entrances[e.id]; ok {
		return e
	}
	p.exits[e.id] = e
	return nil
}

func (p *Parkinglot) AddFloor(pf *ParkingFloor) error {
	pf.globalDisplayBoard = p.globalDisplayBoard
	pf.floorDisplayBoard = NewDisplayBoard()
	if _, ok := p.entrances[pf.id]; ok {
		return pf
	}
	p.parkingFloors[pf.id] = pf
	return nil
}

func (p *Parkinglot) AddParkingSpot(flooId string, ps *ParkingSpot) error {
	if val, ok := p.parkingFloors[flooId]; ok {
		return val.addParkingSpot(ps)
	} else {
		return val
	}
}

func (p Parkinglot) GetGlobalDisplayBoard() *ParkingDisplayBoard {
	return p.globalDisplayBoard
}

func (p Parkinglot) GetExits(exitID string) *ExitPanel {
	return p.exits[exitID]
}

func (p Parkinglot) GetEntries(entranceID string) *EntrancePanel {
	return p.entrances[entranceID]
}

func (p Parkinglot) GetParkingFloor(parkingFloorID string) *ParkingFloor {
	return p.parkingFloors[parkingFloorID]
}

func (p Parkinglot) GetParkingFloorMap() map[string]*ParkingFloor {
	return p.parkingFloors
}
