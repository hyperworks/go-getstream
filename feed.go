package getstream

// FeedID is a typealias of string to create some value safety
type FeedID string

// Feed is the interface bundling all Feed Types
// It exposes methods needed for all Types
type Feed interface {
	Signature() string
	FeedID() FeedID
	Token() string
	SignFeed(signer *Signer)
	GenerateToken(signer *Signer) string
}

// GeneralFeed is a container for Feeds returned from request
// The specific Type will be unknown so no Actions are associated with a GeneralFeed
type GeneralFeed struct {
	Client   *Client
	FeedSlug string
	UserID   string
	token    string
}

// Signature is used to sign Requests : "FeedSlugUserID Token"
func (f *GeneralFeed) Signature() string {
	return f.FeedSlug + f.UserID + " " + f.Token()
}

// FeedID is the combo if the FeedSlug and UserID : "FeedSlug:UserID"
func (f *GeneralFeed) FeedID() FeedID {
	return FeedID(f.FeedSlug + ":" + f.UserID)
}

// SignFeed sets the token on a Feed
func (f *GeneralFeed) SignFeed(signer *Signer) {
	f.token = signer.generateToken(f.FeedSlug + f.UserID)
}

// Token returns the token of a Feed
func (f *GeneralFeed) Token() string {
	return f.token
}

// GenerateToken returns a new Token for a Feed without setting it to the Feed
func (f *GeneralFeed) GenerateToken(signer *Signer) string {
	return signer.generateToken(f.FeedSlug + f.UserID)
}
