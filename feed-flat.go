package getstream

import (
	"encoding/json"
	"errors"
	"regexp"
)

// FlatFeed is a getstream FlatFeed
// Use it to for CRUD on FlatFeed Groups
type FlatFeed struct {
	Client   *Client
	FeedSlug string
	UserID   string
	token    string
}

// Signature is used to sign Requests : "FeedSlugUserID Token"
func (f *FlatFeed) Signature() string {
	return f.FeedSlug + f.UserID + " " + f.Token()
}

// FeedID is the combo if the FeedSlug and UserID : "FeedSlug:UserID"
func (f *FlatFeed) FeedID() FeedID {
	return FeedID(f.FeedSlug + ":" + f.UserID)
}

// SignFeed sets the token on a Feed
func (f *FlatFeed) SignFeed(signer *Signer) {
	f.token = signer.generateToken(f.FeedSlug + f.UserID)
}

// Token returns the token of a Feed
func (f *FlatFeed) Token() string {
	return f.token
}

// GenerateToken returns a new Token for a Feed without setting it to the Feed
func (f *FlatFeed) GenerateToken(signer *Signer) string {
	return signer.generateToken(f.FeedSlug + f.UserID)
}

// AddActivity is Used to post an Activity to a FlatFeed
func (f *FlatFeed) AddActivity(activity *FlatFeedActivity) (*FlatFeedActivity, error) {

	input, err := activity.Input()
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(input)
	if err != nil {
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

// Activities returns a list of Activities for a FlatFeedGroup
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

// RemoveActivity removes an Activity from a FlatFeedGroup
func (f *FlatFeed) RemoveActivity(input *FlatFeedActivity) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + input.ID + "/"

	return f.del(endpoint, f.Signature(), nil)
}

// RemoveActivityByForeignID removes an Activity from a FlatFeedGroup by ForeignID
func (f *FlatFeed) RemoveActivityByForeignID(input *FlatFeedActivity) error {

	r, err := regexp.Compile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")
	if err != nil {
		return err
	}
	if !r.MatchString(input.ForeignID) {
		return errors.New("invalid ForeignID")
	}

	if input.ForeignID == "" {
		return errors.New("no ForeignID")
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + input.ForeignID + "/"

	payload, err := json.Marshal(map[string]string{
		"foreign_id": "1",
	})
	if err != nil {
		return err
	}

	return f.del(endpoint, f.Signature(), payload)
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
