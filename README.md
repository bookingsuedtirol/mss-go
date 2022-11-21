# mss-go

MSS API client for Go projects

## Dependencies

This library requires an IANA Time Zone database to be present on the operating system (because it uses https://pkg.go.dev/time#LoadLocation). This database comes pre-installed with many Linux distros. If it’s unavailable (such as on Windows), https://pkg.go.dev/time/tzdata can be imported in the main program (which uses mss-go) instead.

## Available methods

- [x] getHotelList ✓
- [x] getSpecialList ✓
- [x] getRoomList ✓
- [x] getPriceList ✓
- [x] getRoomAvailability ✓
- [x] prepareBooking ✓
- [x] getBooking ✓
- [x] cancelBooking ✓
- [x] createInquiry
- [x] getUserSources
- [x] getLocationList ✓
- [x] getMasterpackagesList
- [x] getThemeList
- [x] validateCoupon ✓

Warning: Only the methods with a ✓ next to them have been tested so far.

## Before running examples/tests

Set the environment variables with:

```Bash
export $(make env)
```

## Examples

Run `make simple` or `make advanced`

## Tests

Run `make test`
