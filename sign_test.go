package getstream_test

import (
	"testing"
	. "github.com/hyperworks/go-getstream"
	a "github.com/stretchr/testify/assert"
)

const (
	TestAPIKey    = "ufx8newsg6ws"
	TestAPISecret = "37vnzuw6w6h39g2d6vvrkxvyvud5wx55dz3mn5vjvqzd59yy9dwxdzbkabv6crf6"
	TestFeedID    = "test-feed-1"
	TestToken     = "YjvT7Rwj-iqO2wlJfWtUVRts18M"
)

func TestSign(t *testing.T) {
	result := Sign(TestAPISecret, TestFeedID)
	a.Equal(t, TestToken, result)
}
