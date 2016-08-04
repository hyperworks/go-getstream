package getstream

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FlatFeed struct {
	Client   *Client
	FeedSlug string
	UserID   string
	token    string
}

func (f *FlatFeed) Signature() string {
	return f.FeedSlug + f.UserID + " " + f.Token()
}

func (f *FlatFeed) FeedID() string {
	return f.FeedSlug + ":" + f.UserID
}

func (f *FlatFeed) SignFeed(signer *Signer) {
	f.token = signer.generateToken(f.FeedSlug + f.UserID)
}

func (f *FlatFeed) Token() string {
	return f.token
}

func (f *FlatFeed) GenerateToken(signer *Signer) string {
	return signer.generateToken(f.FeedSlug + f.UserID)
}

// get request helper
func (f *FlatFeed) get(path string, signature string) ([]byte, error) {
	res, err := f.request("GET", path, signature, nil)
	return res, err
}

// post request helper
func (f *FlatFeed) post(path string, signature string, payload []byte) ([]byte, error) {
	res, err := f.request("POST", path, signature, payload)
	return res, err
}

// delete request helper
func (f *FlatFeed) del(path string, signature string) error {
	_, err := f.request("DELETE", path, signature, nil)
	return err
}

// request helper
func (f *FlatFeed) request(method, path string, signature string, payload []byte) ([]byte, error) {

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
	if f.Token() != "" {
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

func (f *FlatFeed) AddActivity(activity *FlatFeedActivity) (*FlatFeedActivity, error) {

	input := activity.Input()

	payload, err := json.Marshal(input)
	if err != nil {
		fmt.Println("marshal input error")
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	resultBytes, err := f.post(endpoint, f.Signature(), payload)
	if err != nil {
		return nil, err
	}

	output := &postFlatFeedOutputActivity{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	return output.Activity(), err
}

// func (f *FlatFeed) AddActivities(input []*PostFlatFeedInput) error {
// 	signeds := make([]*Activity, len(input), len(input))
// 	for i, activityInput := range input {
// 		signedActivityInput := f.Client.signer.signActivity(*activityInput)
// 		signeds[i] = signedActivityInput.Activity
// 	}
//
// 	_ = signeds
// 	// TODO: A result type to recieve the listing result.
// 	panic("not yet implemented.")
// }

func (f *FlatFeed) Activities(input *GetFlatFeedInput) (*GetFlatFeedOutput, error) {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	result, err := f.get(endpoint, f.Signature())
	if err != nil {
		return nil, err
	}

	output := &getFlatFeedOutput{}
	err = json.Unmarshal(result, output)
	if err != nil {
		return nil, err
	}

	return output.Output(), err
}

func (f *FlatFeed) RemoveActivity(id string) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + id + "/"

	return f.del(endpoint, f.Signature())
}

func (f *FlatFeed) Follow(feed, id string) error {
	panic("not implemented.")
}

func (f *FlatFeed) Unfollow(feed, id string) error {
	panic("not implemented.")
}

// func (f *FlatFeed) Followers(opt *Options) ([]*FlatFeed, error) {
// 	panic("not implemented.")
// }
//
// func (f *FlatFeed) url() string {
// 	return "feed/" + f.FlatFeedSlug + "/" + f.UserID + "/"
// }
