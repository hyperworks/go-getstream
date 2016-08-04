package getstream

type Feed interface {
	Signature() string
	FeedID() FeedID
	Token() string
	SignFeed(signer *Signer)
	GenerateToken(signer *Signer) string
}

type GeneralFeed struct {
	Client   *Client
	FeedSlug string
	UserID   string
	token    string
}

func (f *GeneralFeed) Signature() string {
	return f.FeedSlug + f.UserID + " " + f.Token()
}

func (f *GeneralFeed) FeedID() FeedID {
	return FeedID(f.FeedSlug + ":" + f.UserID)
}

func (f *GeneralFeed) SignFeed(signer *Signer) {
	f.token = signer.generateToken(f.FeedSlug + f.UserID)
}

func (f *GeneralFeed) Token() string {
	return f.token
}

func (f *GeneralFeed) GenerateToken(signer *Signer) string {
	return signer.generateToken(f.FeedSlug + f.UserID)
}
