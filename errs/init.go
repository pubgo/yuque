package errs

import "errors"

var ErrAppNotFound = errors.New("application not found")
var ErrAppInvalid = errors.New("application invalid")
// authorization code is invalid
// authorization client is invalid
// grant_type invalid
