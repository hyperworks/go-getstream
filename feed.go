package getstream

import ()

type Feed struct {
	*Client
	slug Slug
}

func (f *Feed) Slug() Slug { return f.slug }

func (f *Feed) AddActivity(activity *Activity) (*Activity, error) {
	activity = SignActivity(f.secret, activity)

	result := &Activity{}
	e := f.post(result, "feed/"+f.slug.Slug+"/"+f.slug.ID+"/", f.slug, activity)
	// TODO: inspect result
	return result, e
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
