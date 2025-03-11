# mss-go

MSS API client for Go projects

## Dependencies

This library relies on the IANA Time Zone Database, which must be available on the operating system. It uses [time.LoadLocation](https://pkg.go.dev/time#LoadLocation) to load time zone data. Many Linux distributions include this database by default, but on Alpine Linux, you must install it explicitly:

```Shell
apk add --no-cache tzdata
```

For Windows support, add the following import to your main Go file:

```Go
import _ "time/tzdata"
```

For details, see [time/tzdata](https://pkg.go.dev/time/tzdata).

## Available methods

- [x] getHotelList ✓
- [x] getSpecialList ✓
- [x] getRoomList ✓
- [x] getPriceList ✓
- [x] getRoomAvailability ✓
- [x] prepareBooking ✓
- [x] getBooking ✓
- [x] cancelBooking ✓
- [x] createInquiry ✓
- [x] getUserSources ✓
- [x] getLocationList ✓
- [x] getMasterpackagesList
- [x] getThemeList
- [x] validateCoupon ✓

Warning: Only the methods with a ✓ next to them have been tested so far.

## Before running examples/tests

Set the environment variables with:

```Shell
export $(make env)
```

## Examples

Run `make simple` or `make advanced`

## Tests

Run `make test`
