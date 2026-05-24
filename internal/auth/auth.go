package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the headers of an HTTP request
// Example:
// Authorization: ApiKey {key value here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("authorization")
	if val == "" {
		return "", errors.New("no [Authorization] Header provided")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed [Authorization] Header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed frst part of [Authorization] Header")
	}
	return vals[1], nil
}
