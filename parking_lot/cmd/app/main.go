package main

import (
	"lld-parking-lot/account"
	"lld-parking-lot/parking"
	"lld-parking-lot/vehicle"
)

func main() {
	admin := account.NewAdmin("admin", "admin")

	// initialize a parkingfloor
	groundFloor := parking.NewParkingFloor(admin.GetGlobalDisplayBoard())
	admin.AddParkingFloor(groundFloor) // add parkingfloor to parkinglot

	// initialize a parkingspot
	parkingSpot := parking.NewParkingSpot(groundFloor, parking.CarSpot)
	admin.AddParkingSpot(groundFloor.GetFloorId(), parkingSpot) // add parkingspot to floor in parkinglot

	// initialize an entrance panel
	groundEntrancePanel := parking.NewEntrancePanel(admin.GetGlobalDisplayBoard(), admin.GetParkingFloorMap())
	admin.AddEntrancePanel(groundEntrancePanel) // add entrancepanel to parkinglot

	// initialize an exit panel
	groundFloorExitPanel := parking.NewExitPanel()
	admin.AddExitPanel(groundFloorExitPanel) // add exit panel to parkinglot

	attendant := account.NewParkingAttendant(groundFloorExitPanel, "attendant", "attendant")

	v := vehicle.Vehicle{
		RegistrationNumber:    "some number",
		VehicleType:           vehicle.Car,
		IsHandicappedRequired: false,
	}

	attendant.AssignTicket(groundEntrancePanel, &v)
	attendant.ProcesssTicket(v.GetTicket())
}
