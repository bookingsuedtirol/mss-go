# mss-go

MSS API client for Go projects

**Important**: Do not use this in production yet!

## Available methods

- [x] getHotelList
- [x] getSpecialList
- [x] getRoomList
- [x] getPriceList
- [x] getRoomAvailability
- [x] prepareBooking
- [x] getBooking
- [x] createInquiry
- [x] getUserSources
- [x] getLocationList
- [x] getMasterpackagesList
- [x] getThemeList
- [x] getSeoTexts
- [x] validateCoupon

## Before running examples/tests

Set the environment variables with:

```Bash
export $(grep -v '^#' examples/.env | xargs)
```

## Examples

Run `go run examples/simple/*`, `go run examples/advanced/*` etc.

## Tests

Run `go test -v`
