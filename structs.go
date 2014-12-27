package getstream

import (
	"time"
)

type Activity struct {
	ID        string `json:"id,omitempty"`
	Actor     Slug   `json:"actor"`
	Verb      string `json:"verb"`
	Object    Slug   `json:"object"`
	Target    *Slug  `json:"target,omitempty"`
	RawTime   string `json:"time,omitempty"`
	To        []Slug `json:"to,omitempty"`
	ForeignID string `json:"foreign_id,omitempty"`
}

type Options struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`

	IdGTE string `json:"id_gte"`
	IdGT  string `json:"id_gt"`
	IdLTE string `json:"id_lte"`
	IdLT  string `json:"id_lt"`

	Feeds    []*Feed `json:"feeds"`
	MarkRead bool    `json:"mark_read"`
	MarkSeen bool    `json:"mark_seen"`
}

type Notification struct {
	Data    *Update `json"data"`
	Channel string  `json:"channel"`
}

type Update struct {
	Deletes []*Activity
	Inserts []*Activity

	UnreadCount int
	UnseenCount int
	PublishedAt time.Time
}
