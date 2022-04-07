package hotelfeature

import "github.com/HGV/mss-go/shared"

const (
	Garage shared.HotelFeature = 1 << iota
	ElevatorLift
	Restaurant
	Gym
	Wellness
	SpaCuisineHealthFoods
	ContinentalBreakfastLuncheon
	BreakfastBuffet
	OutdoorPool
	IndoorPool
	Bar
	BarrierFree
	Wlan
	ShuttleService
	Childcare
	SmallPetsAllowed
	BeautyFarm
	CentralLocation shared.HotelFeature = 2 << iota
	CoveredParking
	OpenParking
	Massages
	Sauna
	SteamBath
	PublicBar
	DogsAllowed
)
