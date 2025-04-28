package sdk

var (
	_globalClient IClient
)

func GlobalClient() IClient {
	return _globalClient
}

func SetGlobalClient(c IClient) {
	_globalClient = c
}
