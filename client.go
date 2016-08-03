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

	key      string
	secret   string
	appID    string
	location string // https://location-api.getstream.io/api/
}

// New returns a getstream client.
// Params :
// - api key
// - api secret
// - appID
// - region
func New(key, secret, appID, location string) *Client {
	baseURLStr := "https://api.getstream.io/api/v1.0/"
	if location != "" {
		baseURLStr = "https://" + location + "-api.getstream.io/api/v1.0/"
	}

	baseURL, e := url.Parse(baseURLStr)
	if e != nil {
		panic(e) // failfast, url shouldn't be invalid anyway.
	}

	return &Client{
		http:    &http.Client{},
		baseURL: baseURL,

		key:      key,
		secret:   secret,
		appID:    appID,
		location: location,
	}
}

// BaseURL returns the getstream URL for your location
func (c *Client) BaseURL() *url.URL { return c.baseURL }

// Feed returns a getstream feed
// Slug is the FeedGroup name
// id is the Specific Feed inside a FeedGroup
// to get the feed for Bob you would pass something like "user" as slug and "bob" as the id
func (c *Client) Feed(slug, id string) *Feed {
	return &Feed{
		Client: c,
		slug:   SignSlug(c.secret, Slug{slug, id, ""}),
	}
}

// get request helper
func (c *Client) get(path string, slug Slug) ([]byte, error) {
	res, err := c.request("GET", path, slug, nil)
	return res, err
}

// post request helper
func (c *Client) post(path string, slug Slug, payload []byte) ([]byte, error) {
	res, err := c.request("POST", path, slug, payload)
	return res, err
}

// delete request helper
func (c *Client) del(path string, slug Slug) error {
	_, err := c.request("DELETE", path, slug, nil)
	return err
}

// request helper
func (c *Client) request(method, path string, slug Slug, payload []byte) ([]byte, error) {

	// create url.URL instance with query params
	absURL, err := c.absoluteUrl(path)
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
	if slug.Token != "" {
		req.Header.Set("Authorization", slug.Signature())
	}

	// perform the http request
	resp, err := c.http.Do(req)
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
		if err = json.Unmarshal(respErr, err); err != nil {
			return nil, errors.New(string(respErr))
		}
		return nil, err
	}
}

// absoluteUrl create a url.URL instance and sets query params (bad bad bad!!!)
func (c *Client) absoluteUrl(path string) (result *url.URL, e error) {

	fmt.Println("fix me, I'm a mutant")

	if result, e = url.Parse(path); e != nil {
		return nil, e
	}

	// DEBUG: Use this line to send stuff to a proxy instead.
	// c.baseURL, _ = url.Parse("http://0.0.0.0:8000/")
	result = c.baseURL.ResolveReference(result)

	qs := result.Query()
	qs.Set("api_key", c.key)
	if c.location == "" {
		qs.Set("location", "unspecified")
	} else {
		qs.Set("location", c.location)
	}
	result.RawQuery = qs.Encode()

	return result, nil
}
