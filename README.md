# mss-go

MSS API client for Go projects

**Important**: Do not use this in production yet!

## Available methods

- [x] getHotelList ✓
- [x] getSpecialList ✓
- [x] getRoomList
- [x] getPriceList
- [x] getRoomAvailability ✓
- [x] prepareBooking
- [x] getBooking ✓
- [x] cancelBooking ✓
- [x] createInquiry
- [x] getUserSources
- [x] getLocationList ✓
- [x] getMasterpackagesList
- [x] getThemeList
- [x] validateCoupon

Warning: Only the methods with a check mark next to them have been tested so far.

## Before running examples/tests

Set the environment variables with:

```Bash
export $(make env)
```

## Examples

Run `make simple` or `make advanced`

## Tests

Run `make test`
