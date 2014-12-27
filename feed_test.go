package getstream_test

import (
	"testing"
	// . "github.com/hyperworks/go-getstream"
	a "github.com/stretchr/testify/assert"
)

func TestFeed_AddActivity(t *testing.T) {
	t.Skip()

	client := ConnectTestClient("")
	feed := client.Feed(TestFeedSlug.Slug, TestFeedSlug.ID)
	activity := NewTestActivity()

	result, e := feed.AddActivity(activity)
	a.NoError(t, e)
	a.NotEqual(t, activity, result, "AddActivity should not modify existing instance.")
	a.NotNil(t, result)
	a.NotEmpty(t, result.ID)
}
