package hotel_type

import "github.com/HGV/mss-go/shared"

const (
	Hotel shared.HotelType = 1 << iota
	SkiSchool
	Residence
	BBAndAppartmentsPriv shared.HotelType = 2 << iota
	FarmVacation
	MountainInn
	CampingSite
	HolidayHome
	YouthHostel
	Guesthouse
	Refuge
	Garni
	Inn
)
