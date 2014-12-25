package getstream

import (
	"time"
)

type Activity struct {
	Actor     string    `json:"actor"`
	Verb      string    `json:"verb"`
	Object    string    `json:"object"`
	Target    string    `json:"target"`
	Time      time.Time `json:"time"`
	To        []string  `json:"to"`
	ForeignID string    `json:"foreign_id"`
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
