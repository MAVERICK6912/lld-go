package customerrors

import "errors"

var (
	ErrorParkingSpotUnavailable       error = errors.New("no parking spot available")
	ErrorTicketAlreadyAssigned        error = errors.New("ticket is already assigned to this vehicle")
	ErrorParkingSpotDoesNotExist      error = errors.New("parking spot does not exist")
	ErrorVehicleAlreadyAssignedToSpot error = errors.New("vehicle is already assigned to spot")
	ErrorVehicleNotAssignedToSpot     error = errors.New("no vehicle assigned to spot")
	ErrorInvalidVehicleType           error = errors.New("invalid vehicle type")
)
