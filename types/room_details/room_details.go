package room_details

import "github.com/HGV/mss-go/request"

const (
	BasicInfo             request.RoomDetails = 4
	Title                 request.RoomDetails = 8
	RoomImages            request.RoomDetails = 32
	RoomFacilitiesFilter  request.RoomDetails = 64
	RoomDescription       request.RoomDetails = 256
	RoomFacilitiesDetails request.RoomDetails = 4096
	RoomFeatures          request.RoomDetails = 32768
	RoomNumbers           request.RoomDetails = 65536
)
