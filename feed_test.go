package getstream_test

import (
	"testing"
	// . "github.com/hyperworks/go-getstream"
	a "github.com/stretchr/testify/assert"
)

func TestFeed(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	client := ConnectTestClient("")
	feed := client.Feed(TestFeedSlug.Slug, TestFeedSlug.ID)
	activity := NewTestActivity()

	t.Log("adding activity...")
	addedActivity, e := feed.AddActivity(activity)
	a.NoError(t, e)
	a.NotEqual(t, activity, addedActivity, "AddActivity should not modify existing instance.")
	a.NotNil(t, addedActivity)
	a.NotEmpty(t, addedActivity.ID)

	t.Log("listing added activities...")
	activities, e := feed.Activities(nil)
	a.NoError(t, e)
	a.NotEmpty(t, activities)
	a.Len(t, activities, 1) // otherwise we might be getting result from another test run.
	a.Equal(t, addedActivity.ID, activities[0].ID)

	t.Log("removing added activity...")
	e = feed.RemoveActivity(addedActivity.ID)
	a.NoError(t, e)

	t.Log("listing added activities again...")
	activities, e = feed.Activities(nil)
	a.NoError(t, e)
	a.Empty(t, activities)
}
