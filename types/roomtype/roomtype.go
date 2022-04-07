package roomtype

import "github.com/HGV/mss-go/shared"

const (
	Room shared.RoomType = 1 << iota
	Apartment
	CampingPitch
)
