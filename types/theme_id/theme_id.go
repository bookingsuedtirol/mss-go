package theme_id

import "github.com/HGV/mss-go/request"

const (
	Hiking request.ThemeID = iota + 1
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
	BikeHotelsMountainbike request.ThemeID = iota + 2
	BikeHotelsBikeTours
	BikeHotelsRacingBike
	FamilyHotels
	FamilyHotelsNatureDetective
	FamilyHotel request.ThemeID = iota + 3
	NatureDetectivSummer
	NatureDetectivWinter
)
