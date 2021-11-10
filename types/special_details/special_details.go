package special_details

import "github.com/HGV/mss-go/request"

const (
	BasicInfo request.SpecialDetails = 1 << iota
	Title
	Descriptions
	Seasons
	Images
	Themes
	IncludedServices
	HotelIncludedServices
	HotelMandatoryServices
)
