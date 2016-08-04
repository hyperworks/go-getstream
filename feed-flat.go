package getstream

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
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

	input, err := activity.input()
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

// AddActivities is Used to post multiple Activities to a FlatFeed
func (f *FlatFeed) AddActivities(activities []*FlatFeedActivity) ([]*FlatFeedActivity, error) {

	var inputs []*postFlatFeedInputActivity

	for _, activity := range activities {
		activity, err := activity.input()
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, activity)
	}

	payload, err := json.Marshal(map[string][]*postFlatFeedInputActivity{
		"activities": inputs,
	})
	if err != nil {
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	resultBytes, err := f.post(endpoint, f.Signature(), payload)
	if err != nil {
		return nil, err
	}

	output := &postFlatFeedOutputActivities{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	var outputActivities []*FlatFeedActivity
	for _, outputActivity := range output.Activities {
		activity := outputActivity.Activity()
		outputActivities = append(outputActivities, activity)
	}

	return outputActivities, err
}

// Activities returns a list of Activities for a FlatFeedGroup
func (f *FlatFeed) Activities(input *GetFlatFeedInput) (*GetFlatFeedOutput, error) {

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/"

	result, err := f.get(endpoint, f.Signature(), payload)
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

// FollowFeedWithCopyLimit sets a Feed to follow another target Feed
// CopyLimit is the maximum number of Activities to Copy from History
func (f *FlatFeed) FollowFeedWithCopyLimit(target *FlatFeed, copyLimit int) error {
	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/"

	input := postFlatFeedFollowingInput{
		Target:            string(target.FeedID()),
		ActivityCopyLimit: copyLimit,
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}

	_, err = f.post(endpoint, f.Signature(), payload)
	return err

}

// Unfollow is used to Unfollow a target Feed
func (f *FlatFeed) Unfollow(target *FlatFeed) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/" + string(target.FeedID()) + "/"

	return f.del(endpoint, f.Signature(), nil)

}

// UnfollowKeepingHistory is used to Unfollow a target Feed while keeping the History
// this means that Activities already visibile will remain
func (f *FlatFeed) UnfollowKeepingHistory(target *FlatFeed) error {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/" + string(target.FeedID()) + "/"

	payload, err := json.Marshal(map[string]string{
		"keep_history": "1",
	})
	if err != nil {
		return err
	}

	return f.del(endpoint, f.Signature(), payload)

}

// FollowersWithLimitAndSkip returns a list of GeneralFeed following the current FlatFeed
func (f *FlatFeed) FollowersWithLimitAndSkip(limit int, skip int) ([]*GeneralFeed, error) {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "followers" + "/"

	payload, err := json.Marshal(&getFlatFeedFollowersInput{
		Limit: limit,
		Skip:  skip,
	})
	if err != nil {
		return nil, err
	}

	resultBytes, err := f.get(endpoint, f.Signature(), payload)

	output := &getFlatFeedFollowersOutput{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	var outputFeeds []*GeneralFeed
	for _, result := range output.Results {

		feed := GeneralFeed{}

		match, err := regexp.MatchString(`^.*?:.*?$`, result.FeedID)
		if err != nil {
			continue
		}

		if match {
			firstSplit := strings.Split(result.FeedID, ":")

			feed.FeedSlug = firstSplit[0]
			feed.UserID = firstSplit[1]
		}

		outputFeeds = append(outputFeeds, &feed)
	}

	return outputFeeds, err

}

// FollowingWithLimitAndSkip returns a list of GeneralFeed followed by the current FlatFeed
func (f *FlatFeed) FollowingWithLimitAndSkip(limit int, skip int) ([]*GeneralFeed, error) {

	endpoint := "feed/" + f.FeedSlug + "/" + f.UserID + "/" + "following" + "/"

	payload, err := json.Marshal(&getFlatFeedFollowersInput{
		Limit: limit,
		Skip:  skip,
	})
	if err != nil {
		return nil, err
	}

	resultBytes, err := f.get(endpoint, f.Signature(), payload)

	output := &getFlatFeedFollowersOutput{}
	err = json.Unmarshal(resultBytes, output)
	if err != nil {
		return nil, err
	}

	var outputFeeds []*GeneralFeed
	for _, result := range output.Results {

		feed := GeneralFeed{}

		match, err := regexp.MatchString(`^.*?:.*?$`, result.FeedID)
		if err != nil {
			continue
		}

		if match {
			firstSplit := strings.Split(result.TargetID, ":")

			feed.FeedSlug = firstSplit[0]
			feed.UserID = firstSplit[1]
		}

		outputFeeds = append(outputFeeds, &feed)
	}

	return outputFeeds, err

}
