package getstream

import (
	"encoding/json"
	"fmt"
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

func (f *Feed) AddActivity(input *PostActivityInput) (*PostActivityOutput, error) {

	signedActivityInput := f.Client.signer.signActivity(*input)

	payload, err := json.Marshal(signedActivityInput.Activity)
	if err != nil {
		fmt.Println("marshal input error")
		return nil, err
	}

	resultBytes, err := f.post(f.url(), f.Signature(), payload)
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

	result, err := f.get(f.url(), f.Signature())
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
	return f.del(f.url()+id+"/", f.Signature())
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

func (f *Feed) url() string {
	return "feed/" + f.FeedSlug + "/" + f.UserID + "/"
}
