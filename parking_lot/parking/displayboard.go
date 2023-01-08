package parking

import (
	"fmt"

	"github.com/google/uuid"
)

type ParkingDisplayBoard struct {
	id               string
	parkingSpotCount map[ParkingSpotType]int
}

func NewDisplayBoard() *ParkingDisplayBoard {
	return &ParkingDisplayBoard{
		id:               uuid.NewString(),
		parkingSpotCount: make(map[ParkingSpotType]int),
	}
}

func (pdb *ParkingDisplayBoard) Show() {
	for k, v := range pdb.parkingSpotCount {
		fmt.Printf("%d available for %s parking spot type.", v, k)
	}
}

func (pdb *ParkingDisplayBoard) ChangeCount(pst *ParkingSpotType, count int) {
	pdb.parkingSpotCount[*pst] += count
}
