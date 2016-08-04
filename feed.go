package getstream

type Feed interface {
	Signature() string
	FeedID() string
	Token() string
	SignFeed(signer *Signer)
	GenerateToken(signer *Signer) string
}
