package getstream

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	http    *http.Client
	baseURL *url.URL // https://api.getstream.io/api/

	key      string
	secret   string
	appID    string
	location string // https://location-api.getstream.io/api/
}

func Connect(key, secret, appID, location string) *Client {
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

func (c *Client) BaseURL() *url.URL { return c.baseURL }

func (c *Client) Feed(slug, id string) *Feed {
	return &Feed{
		Client: c,
		slug:   SignSlug(c.secret, Slug{slug, id, ""}),
	}
}

func (c *Client) post(result interface{}, path string, slug Slug, payload interface{}) error {
	absUrl, e := c.absoluteUrl(path)
	if e != nil {
		return e
	}

	buffer, e := json.Marshal(payload)
	if e != nil {
		return e
	}

	req, e := http.NewRequest("POST", absUrl.String(), bytes.NewBuffer(buffer))
	if e != nil {
		return e
	}

	req.Header.Set("Content-Type", "application/json")
	if slug.Token != "" {
		req.Header.Set("Authorization", slug.Signature())
	}

	resp, e := c.http.Do(req)
	if e != nil {
		return e
	}

	buffer, e = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e != nil {
		return e
	}

	switch {
	case 200 <= resp.StatusCode && resp.StatusCode < 300: // SUCCESS
		if e = json.Unmarshal(buffer, result); e != nil {
			return e
		}

	default:
		err := &Error{}
		if e = json.Unmarshal(buffer, err); e != nil {
			panic(e)
			return errors.New(string(buffer))
		}

		return err
	}

	return nil
}

func (c *Client) absoluteUrl(path string) (result *url.URL, e error) {
	if result, e = url.Parse(path); e != nil {
		return nil, e
	}

	// DEBUG: Use this line to send stuff to a proxy instead.
	// c.baseURL, _ = url.Parse("http://0.0.0.0:8080/")
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
