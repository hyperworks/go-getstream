package getstream

import (
	"time"
)

/*{"code": 5, "detail": "Please use signedto instead of signedTo for your field name", "
duration": "36ms", "exception": "CustomFieldException", "status_code": 400}           */
type Error struct {
	Code       int `json:"code"`
	StatusCode int `json:"status_code"`

	Detail      string `json:"detail"`
	RawDuration string `json:"duration"`
	Exception   string `json:"exception"`
}

var _ error = &Error{}

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
