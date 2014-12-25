package getstream

type Feed struct {
	name string
	id   string
}

func (f *Feed) AddActivity(activity *Activity) error {
	panic("not implemented.")
}

func (f *Feed) AddActivities(activities []*Activity) error {
	panic("not implemented.")
}

func (f *Feed) Activities(opt *Options) ([]*Activity, error) {
	panic("not implemented.")
}

func (f *Feed) Follow(feed, id string) error {
	panic("not implemented.")
}

func (f *Feed) Unfollow(feed, id string) error {
	panic("not implemented.")
}

func (f *Feed) Followers(opt *Options) ([]*Feed, error) {
	panic("not implemented.")
}
