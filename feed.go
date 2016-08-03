package getstream

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Feed struct {
	Client   *Client
	FeedSlug string
	UserID   string
	Token    string
}

func (f *Feed) Signature() string {
	return f.FeedSlug + f.UserID + " " + f.Token
}

func (f *Feed) FeedID() string {
	return f.FeedSlug + ":" + f.UserID
}

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

	// debug Println
	//fmt.Println(string(body))

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
			return nil, err
		}
		return nil, errors.New(string(respErr))
	}
}

func (f *Feed) AddActivity(input *PostActivityInput) (*PostActivityOutput, error) {

	signedActivityInput := f.Client.signer.signActivity(*input)

	payload, err := json.Marshal(signedActivityInput.Activity)
	if err != nil {
		fmt.Println("marshal input error")
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	resultBytes, err := f.post(endpoint, f.Signature(), payload)
	if err != nil {
		return nil, err
	}

	output := &PostActivityOutput{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (f *Feed) AddActivities(input []*PostActivityInput) error {
	signeds := make([]*Activity, len(input), len(input))
	for i, activityInput := range input {
		signedActivityInput := f.Client.signer.signActivity(*activityInput)
		signeds[i] = signedActivityInput.Activity
	}

	_ = signeds
	// TODO: A result type to recieve the listing result.
	panic("not yet implemented.")
}

func (f *Feed) Activities(input *GetActivityInput) (*GetActivityOutput, error) {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	result, err := f.get(endpoint, f.Signature())
	if err != nil {
		return nil, err
	}

	output := &GetActivityOutput{}
	err = json.Unmarshal(result, output)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (f *Feed) RemoveActivity(id string) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + id + "/"

	return f.del(endpoint, f.Signature())
}

func (f *Feed) Follow(feed, id string) error {
	panic("not implemented.")
}

func (f *Feed) Unfollow(feed, id string) error {
	panic("not implemented.")
}

// func (f *Feed) Followers(opt *Options) ([]*Feed, error) {
// 	panic("not implemented.")
// }
//
// func (f *Feed) url() string {
// 	return "feed/" + f.FeedSlug + "/" + f.UserID + "/"
// }
