package getstream

import (
	"encoding/json"
	"regexp"
	"strings"
)

type FlatFeedActivity struct {
	ID      string
	Actor   string
	Verb    string
	Object  string
	Target  string
	RawTime string

	ForeignID string
	Data      json.RawMessage

	To []Feed
}

func (a FlatFeedActivity) Input() *postFlatFeedInputActivity {

	input := postFlatFeedInputActivity{
		ID:        a.ID,
		Actor:     a.Actor,
		Verb:      a.Verb,
		Object:    a.Object,
		Target:    a.Target,
		RawTime:   a.RawTime,
		ForeignID: a.ForeignID,
		Data:      a.Data,
	}

	input.To = []string{}

	for _, feed := range a.To {
		to := feed.FeedID()
		if feed.Token() != "" {
			to += " " + feed.Token()
		}
		input.To = append(input.To, to)
	}

	return &input
}

type postFlatFeedInputActivity struct {
	ID        string          `json:"id,omitempty"`
	Actor     string          `json:"actor"`
	Verb      string          `json:"verb"`
	Object    string          `json:"object"`
	Target    string          `json:"target,omitempty"`
	RawTime   string          `json:"time,omitempty"`
	ForeignID string          `json:"foreign_id,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
	To        []string        `json:"to,omitempty"`
}

type postFlatFeedOutputActivity struct {
	ID        string          `json:"id,omitempty"`
	Actor     string          `json:"actor"`
	Verb      string          `json:"verb"`
	Object    string          `json:"object"`
	Target    string          `json:"target,omitempty"`
	RawTime   string          `json:"time,omitempty"`
	ForeignID string          `json:"foreign_id,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
	To        [][]string      `json:"to,omitempty"`
}

func (a postFlatFeedOutputActivity) Activity() *FlatFeedActivity {

	activity := FlatFeedActivity{
		ID:        a.ID,
		Actor:     a.Actor,
		Verb:      a.Verb,
		Object:    a.Object,
		Target:    a.Target,
		RawTime:   a.RawTime,
		ForeignID: a.ForeignID,
		Data:      a.Data,
	}

	for _, slice := range a.To {
		for _, to := range slice {

			feed := GeneralFeed{}

			match, err := regexp.MatchString(`^.*?:.*? .*?$`, to)
			if err != nil {
				continue
			}

			if match {
				firstSplit := strings.Split(to, ":")
				secondSplit := strings.Split(firstSplit[1], " ")

				feed.FeedSlug = firstSplit[0]
				feed.UserID = secondSplit[0]
				feed.token = secondSplit[1]
			}

			activity.To = append(activity.To, &feed)
		}
	}
	return &activity
}

type GetFlatFeedInput struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	IDGTE int `json:"id_gte,omitempty"`
	IDGT  int `json:"id_gt,omitempty"`
	IDLTE int `json:"id_lte,omitempty"`
	IDLT  int `json:"id_lt,omitempty"`

	Ranking string `json:"ranking,omitempty"`
}

type GetFlatFeedOutput struct {
	Duration   string
	Next       string
	Activities []*FlatFeedActivity
}

type getFlatFeedOutput struct {
	Duration   string                       `json:"duration"`
	Next       string                       `json:"next"`
	Activities []*getFlatFeedOutputActivity `json:"results"`
}

func (a getFlatFeedOutput) Output() *GetFlatFeedOutput {

	output := GetFlatFeedOutput{
		Duration: a.Duration,
		Next:     a.Next,
	}

	for _, activity := range a.Activities {
		output.Activities = append(output.Activities, activity.Activity())
	}

	return &output
}

type getFlatFeedOutputActivity struct {
	ID        string          `json:"id,omitempty"`
	Actor     string          `json:"actor"`
	Verb      string          `json:"verb"`
	Object    string          `json:"object"`
	Target    string          `json:"target,omitempty"`
	RawTime   string          `json:"time,omitempty"`
	To        []string        `json:"to,omitempty"`
	ForeignID string          `json:"foreign_id,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

func (a getFlatFeedOutputActivity) Activity() *FlatFeedActivity {

	activity := FlatFeedActivity{
		ID:        a.ID,
		Actor:     a.Actor,
		Verb:      a.Verb,
		Object:    a.Object,
		Target:    a.Target,
		RawTime:   a.RawTime,
		ForeignID: a.ForeignID,
		Data:      a.Data,
	}

	for _, to := range a.To {

		feed := GeneralFeed{}

		match, err := regexp.MatchString(`^.*?:.*? .*?$`, to)
		if err != nil {
			continue
		}

		if match {
			firstSplit := strings.Split(to, ":")
			secondSplit := strings.Split(firstSplit[1], " ")

			feed.FeedSlug = firstSplit[0]
			feed.UserID = secondSplit[0]
			feed.token = secondSplit[1]
		}

		activity.To = append(activity.To, &feed)
	}
	return &activity
}
