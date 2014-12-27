package getstream_test

import (
	. "github.com/hyperworks/go-getstream"
)

const (
	TestAPIKey           = "ufx8newsg6ws"
	TestAPISecret        = "37vnzuw6w6h39g2d6vvrkxvyvud5wx55dz3mn5vjvqzd59yy9dwxdzbkabv6crf6"
	TestAppID            = "1649"
	TestToken            = "x_oOp5aDI9IAbIjYQw4uFscKi2E"
	TestTargetFeedToken  = "O9kSEnFPFAzoR071IEhYuLhP8mM"
	TestTargetFeedToken2 = "K2S9epudS62ll2PA-BZbS9lECBI"

	TestFeedSignature        = "userflat14483198-3e43-4a91-a2ed-fc88dcf2fd7b " + TestToken
	TestTargetFeedSignature  = "userflat72d8ee6c-c27c-49a5-9311-c5e0e67356e5 " + TestTargetFeedToken
	TestTargetFeedSignature2 = "userflate1f8917c-e6dd-4d06-b6ff-59805d8e2b96 " + TestTargetFeedToken2
)

var (
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
