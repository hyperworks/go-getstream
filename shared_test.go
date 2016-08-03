package getstream_test

import (
	"os"

	. "github.com/hyperworks/go-getstream"
)

const (
	TestToken            = "xfHQwFY2MreP7bTGsoxQ7s2ebLA"
	TestTargetFeedToken  = "3ZIhnU1vw524lXGsOd7wb0DkmrU"
	TestTargetFeedToken2 = "IKjr7WED0O3ROLZGEIFrCSwBo4Y"

	TestFlatFeedSignature    = "userflat14483198-3e43-4a91-a2ed-fc88dcf2fd7b " + TestToken
	TestTargetFeedSignature  = "userflat72d8ee6c-c27c-49a5-9311-c5e0e67356e5 " + TestTargetFeedToken
	TestTargetFeedSignature2 = "userflate1f8917c-e6dd-4d06-b6ff-59805d8e2b96 " + TestTargetFeedToken2
)

var (
	TestAPIKey    = os.Getenv("GETSTREAM_KEY")
	TestAPISecret = os.Getenv("GETSTREAM_SECRET")
	TestAppID     = os.Getenv("GETSTREAM_APPID")

	TestFeedSlug        = Slug{"userflat", "14483198-3e43-4a91-a2ed-fc88dcf2fd7b", ""}
	TestTargetFeedSlug  = Slug{"userflat", "72d8ee6c-c27c-49a5-9311-c5e0e67356e5", ""}
	TestTargetFeedSlug2 = Slug{"userflat", "e1f8917c-e6dd-4d06-b6ff-59805d8e2b96", ""}
	TestObjectSlug      = Slug{"testobject", "2109704a-e048-4f5c-b534-1ff8322b8ae9", ""}
)

func ConnectTestClient(region string) *Client {
	return Connect(TestAPIKey, TestAPISecret, TestAppID, region)
}

func NewTestActivity() *Activity {
	return &Activity{
		Actor:  TestFeedSlug,
		Verb:   "comment",
		Object: TestObjectSlug,
		To:     []Slug{TestTargetFeedSlug, TestTargetFeedSlug2},
	}
}
