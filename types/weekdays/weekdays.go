package weekdays

import "github.com/HGV/mss-go/response"

const (
	Monday response.Weekdays = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
