package themeid

import "github.com/HGV/mss-go/shared"

const (
	Hiking shared.ThemeID = iota + 1
	CyclingMountainbike
	Family
	WellnessHealth
	FoodAndDrink
	Golf
	Culture
	Motorsport
	CarFreeHolidays
	SkiSnowboard
	SummerActivities
	Events
	ChristmasMarkets
	ActiveWinter
	Vitalpina
	VitalpinaBreathe
	BikeHotelsEBike
	BikeHotelsFreeride
	BikeHotelsMountainbike shared.ThemeID = iota + 2
	BikeHotelsBikeTours
	BikeHotelsRacingBike
	FamilyHotels
	FamilyHotelsNatureDetective
	FamilyHotel shared.ThemeID = iota + 3
	NatureDetectivSummer
	NatureDetectivWinter
)
