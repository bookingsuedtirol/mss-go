package error_codes

import "github.com/HGV/mss-go/response"

const (
	GenericError response.ErrorCode = 1 << iota
	AuthenticationError
	InvalidXml
	InvalidMethod
	ResultIdNotInCache
	InvalidMissingParameter
	BookingValidationFailed
	PermissionsDenied
)
