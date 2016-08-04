package getstream

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

// Client is used to connect to getstream.io
type Client struct {
	http    *http.Client
	baseURL *url.URL // https://api.getstream.io/api/

	Key      string
	Secret   string
	AppID    string
	Location string // https://location-api.getstream.io/api/

	signer *Signer
}

// New returns a getstream client.
// Params :
// - api key
// - api secret
// - appID
// - region
func New(key, secret, appID, location string) (*Client, error) {
	baseURLStr := "https://api.getstream.io/api/v1.0/"
	if location != "" {
		baseURLStr = "https://" + location + "-api.getstream.io/api/v1.0/"
	}

	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, err
	}

	return &Client{
		http: &http.Client{
			Timeout: 3 * time.Second,
		},
		baseURL: baseURL,

		Key:      key,
		Secret:   secret,
		AppID:    appID,
		Location: location,

		signer: &Signer{
			Secret: secret,
		},
	}, nil
}

// FlatFeed returns a getstream feed
// Slug is the FlatFeedGroup name
// id is the Specific FlatFeed inside a FlatFeedGroup
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) FlatFeed(feedSlug string, userID string) (*FlatFeed, error) {

	r, err := regexp.Compile(`^\w+$`)
	if err != nil {
		return nil, err
	}
	if !r.MatchString(feedSlug) || !r.MatchString(userID) {
		return nil, errors.New("invalid ForeignID")
	}

	feed := &FlatFeed{
		Client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	feed.SignFeed(c.signer)
	return feed, nil
}

// NotificationFeed returns a getstream feed
// Slug is the NotificationFeedGroup name
// id is the Specific NotificationFeed inside a NotificationFeedGroup
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) NotificationFeed(feedSlug string, userID string) (*NotificationFeed, error) {

	r, err := regexp.Compile(`^\w+$`)
	if err != nil {
		return nil, err
	}
	if !r.MatchString(feedSlug) || !r.MatchString(userID) {
		return nil, errors.New("invalid ForeignID")
	}

	feed := &NotificationFeed{
		Client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	feed.SignFeed(c.signer)
	return feed, nil
}

// BaseURL returns the getstream URL for your location
func (c *Client) BaseURL() *url.URL { return c.baseURL }

// absoluteUrl create a url.URL instance and sets query params (bad!!!)
func (c *Client) absoluteURL(path string) (*url.URL, error) {

	result, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// DEBUG: Use this line to send stuff to a proxy instead.
	// c.baseURL, _ = url.Parse("http://0.0.0.0:8000/")
	result = c.baseURL.ResolveReference(result)

	qs := result.Query()
	qs.Set("api_key", c.Key)
	if c.Location == "" {
		qs.Set("location", "unspecified")
	} else {
		qs.Set("location", c.Location)
	}
	result.RawQuery = qs.Encode()

	return result, nil
}
