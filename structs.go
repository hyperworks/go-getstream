package getstream

type PostActivityInput struct {
	Activity *Activity
	To       []*Feed
}

type PostActivityOutput struct {
	Duration   string      `json:"duration"`
	Next       string      `json:"next"`
	Activities []*Activity `json:"results"`
}

type Activity struct {
	ID        string   `json:"id,omitempty"`
	Actor     string   `json:"actor"`
	Verb      string   `json:"verb"`
	Object    string   `json:"object"`
	Target    string   `json:"target,omitempty"`
	RawTime   string   `json:"time,omitempty"`
	To        []string `json:"to,omitempty"`
	ForeignID string   `json:"foreign_id,omitempty"`
	Data      string   `json:"data,omitempty"`
}

type GetActivityInput struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	IDGTE int `json:"id_gte,omitempty"`
	IDGT  int `json:"id_gt,omitempty"`
	IDLTE int `json:"id_lte,omitempty"`
	IDLT  int `json:"id_lt,omitempty"`

	Ranking string `json:"ranking,omitempty"`
}

type GetActivityOutput struct {
	Duration   string      `json:"duration"`
	Next       string      `json:"next"`
	Activities []*Activity `json:"results"`
}

//
// type Notification struct {
// 	Data    *Update `json"data"`
// 	Channel string  `json:"channel"`
// }
//
// type Update struct {
// 	Deletes []*Activity
// 	Inserts []*Activity
//
// 	UnreadCount int
// 	UnseenCount int
// 	PublishedAt time.Time
// }
