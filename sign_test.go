package getstream_test

import (
	"testing"

	. "github.com/hyperworks/go-getstream"
	a "github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	result := Sign(TestAPISecret, TestFeedSlug.Slug+TestFeedSlug.ID)
	a.Equal(t, TestToken, result)
}

func TestSignSlug(t *testing.T) {
	expected := TestFeedSlug.WithToken(TestToken)
	actual := SignSlug(TestAPISecret, TestFeedSlug)
	a.Equal(t, expected, actual)
	a.Equal(t, TestFlatFeedSignature, actual.Signature())
}

func TestSignActivity(t *testing.T) {
	act := NewTestActivity()
	act = SignActivity(TestAPISecret, act)
	a.Equal(t, TestTargetFeedSignature, act.To[0].Signature())
	a.Equal(t, TestTargetFeedSignature2, act.To[1].Signature())
}
