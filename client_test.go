package getstream_test

import (
	"testing"
	// . "github.com/hyperworks/go-getstream"
	a "github.com/stretchr/testify/assert"
)

func TestClient_BaseURL(t *testing.T) {
	locations := map[string]string{
		"":        "https://api.getstream.io/api/v1.0/",
		"us-east": "https://us-east-api.getstream.io/api/v1.0/",
		"xyz":     "https://xyz-api.getstream.io/api/v1.0/",
	}

	for location, url := range locations {
		client := ConnectTestClient(location)
		a.Equal(t, url, client.BaseURL().String())
	}
}

func TestClient_Feed(t *testing.T) {
	client := ConnectTestClient("")
	feed := client.Feed(TestFeedSlug.Slug, TestFeedSlug.ID)
	a.Equal(t, TestFeedSlug.WithToken(TestToken), feed.Slug())
	a.Equal(t, TestFeedSignature, feed.Slug().Signature())
}
