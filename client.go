package getstream

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
		http:    &http.Client{},
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

// BaseURL returns the getstream URL for your location
func (c *Client) BaseURL() *url.URL { return c.baseURL }

// Feed returns a getstream feed
// Slug is the FeedGroup name
// id is the Specific Feed inside a FeedGroup
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) Feed(feedSlug string, userID string) *Feed {
	feed := &Feed{
		Client:   c,
		FeedSlug: feedSlug,
		UserID:   userID,
	}

	c.signer.signFeed(feed)
	return feed
}

// get request helper
func (f *Feed) get(path string, signature string) ([]byte, error) {
	res, err := f.request("GET", path, signature, nil)
	return res, err
}

// post request helper
func (f *Feed) post(path string, signature string, payload []byte) ([]byte, error) {
	res, err := f.request("POST", path, signature, payload)
	return res, err
}

// delete request helper
func (f *Feed) del(path string, signature string) error {
	_, err := f.request("DELETE", path, signature, nil)
	return err
}

// request helper
func (f *Feed) request(method, path string, signature string, payload []byte) ([]byte, error) {

	// create url.URL instance with query params
	absURL, err := f.Client.absoluteURL(path)
	if err != nil {
		return nil, err
	}

	// create a new http request
	req, err := http.NewRequest(method, absURL.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// set the Auth headers for the http request
	req.Header.Set("Content-Type", "application/json")
	if f.Token != "" {
		req.Header.Set("Authorization", signature)
	}

	// perform the http request
	resp, err := f.Client.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// handle the response
	switch {
	case resp.StatusCode/100 == 2: // SUCCESS
		if body != nil {
			return body, nil
		}
		return nil, nil
	default:
		var respErr []byte
		err = json.Unmarshal(respErr, err)
		if err != nil {
			return nil, errors.New(string(respErr))
		}
		return nil, err
	}
}

// absoluteUrl create a url.URL instance and sets query params (bad bad bad!!!)
func (c *Client) absoluteURL(path string) (*url.URL, error) {

	fmt.Println("fix me, I'm a mutant")

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
