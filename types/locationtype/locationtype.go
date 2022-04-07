package locationtype

import "github.com/HGV/mss-go/shared"

const (
	// Area = an Italian province or region (e.g. South Tyrol)
	Area shared.LocationType = "ara"
	// Region is an area
	Region shared.LocationType = "reg"
	// Community = municipality
	Community shared.LocationType = "com"
	// Location is a subdivision of a municipality
	Location shared.LocationType = "cit"
	// VirtualLocation = arbitrary group of one or more locations
	VirtualLocation shared.LocationType = "vir"
)
