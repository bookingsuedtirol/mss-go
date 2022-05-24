package cancelledstatus

import "github.com/HGV/mss-go/response"

// All statuses can be cancelled with cancelBooking except for Cancelled (1)
const (
	NotCancelled response.CancelledStatus = iota
	Cancelled
	NoShow response.CancelledStatus = iota + 5
	Unknown
)
