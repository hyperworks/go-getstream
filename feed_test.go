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

	addedActivity, e := feed.AddActivity(activity)
	a.NoError(t, e)
	a.NotEqual(t, activity, addedActivity, "AddActivity should not modify existing instance.")
	a.NotNil(t, addedActivity)
	a.NotEmpty(t, addedActivity.ID)
}
