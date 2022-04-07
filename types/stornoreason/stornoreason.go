package stornoreason

import "github.com/HGV/mss-go/request"

const (
	Unknown                       request.StornoReason = 0
	GuestUnavailable              request.StornoReason = 1
	PropertyRequestedCancellation request.StornoReason = 2
	GuestChoseAnotherDestination  request.StornoReason = 3
	GuestChoseAnotherProperty     request.StornoReason = 4
	Other                         request.StornoReason = 99
)
