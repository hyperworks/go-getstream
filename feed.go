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
	e := f.post(result, f.url(), f.slug, activity)
	return result, e
}

func (f *Feed) AddActivities(activities []*Activity) error {
	signeds := make([]*Activity, len(activities), len(activities))
	for i, activity := range activities {
		signeds[i] = SignActivity(f.secret, activity)
	}

	// TODO: A result type to recieve the listing result.
	panic("not yet implemented.")
}

func (f *Feed) Activities(opt *Options) ([]*Activity, error) {
	result := ActivitiesResult{}
	e := f.get(&result, f.url(), f.slug)
	return result.Results, e
}

func (f *Feed) RemoveActivity(id string) error {
	return f.del(f.url()+id+"/", f.slug)
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

func (f *Feed) url() string {
	return "feed/" + f.slug.Slug + "/" + f.slug.ID + "/"
}
