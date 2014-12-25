package getstream

type Client struct {
}

func Connect(key, secret string) (*Client, error) {
	panic("not implemented.")
}

func (c *Client) Feed(name, id string) *Feed {
	return &Feed{name, id}
}
