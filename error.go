package getstream

import (
	"time"
)

// Error is a getstream error
type Error struct {
	Code       int `json:"code"`
	StatusCode int `json:"status_code"`

	Detail      string `json:"detail"`
	RawDuration string `json:"duration"`
	Exception   string `json:"exception"`
}

var _ error = &Error{}

// Duration is the time it took for the request to be handled
func (err *Error) Duration() time.Duration {
	result, e := time.ParseDuration(err.RawDuration)
	if e != nil {
		return time.Duration(0)
	}

	return result
}

func (err *Error) Error() string {
	str := err.Exception
	if err.RawDuration != "" {
		if duration := err.Duration(); duration > 0 {
			str += " (" + duration.String() + ")"
		}
	}

	if err.Detail != "" {
		str += ": " + err.Detail
	}

	return str
}
