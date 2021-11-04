package error_codes

import "github.com/HGV/mss-go/response"

const (
	GenericError response.ErrorCode = 1 << iota
	AuthenticationError
	InvalidXml
	InvalidMethod
	ResultIDNotInCache
	InvalidMissingParameter
	BookingValidationFailed
	PermissionsDenied
)
