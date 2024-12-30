package nix_sdk

var (
	_globalClient *Client
)

func GlobalClient() *Client {
	return _globalClient
}

func SetGlobalClient(c *Client) {
	_globalClient = c
}
