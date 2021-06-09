# mss-go

MSS API client for Go projects

**Important**: Do not use this in production yet!

## Available methods

- [x] getLocationList

To run

1. Export the variables with:
   ```Bash
   export $(grep -v '^#' .env | xargs)
   ```
2. `go run .`

TODO:

- Parse dates as time interface, not as string
- Perhaps parse XML comma list as array?
