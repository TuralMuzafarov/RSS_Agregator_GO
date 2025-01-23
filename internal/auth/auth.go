package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("authentication info is not proper form")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("first parth of the authentication info is incorrect")
	}

	return vals[1], nil
}
