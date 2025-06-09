package openrouter

type OpenrouterConfig struct {
	Host   string
	ApiKey string
	Urls   Urls
	Limit  int
}

type Urls struct {
	apiV1 Endpoints
}

type Endpoints struct {
	key            string
	chatCompletion string
}

var urls = Urls{
	apiV1: Endpoints{
		key:            "/api/v1/key",
		chatCompletion: "/api/v1/chat/completions",
	},
}

var config = OpenrouterConfig{}

func SetConfig(newConfig OpenrouterConfig) {
	config.Host = newConfig.Host
	config.ApiKey = newConfig.ApiKey
	config.Urls = urls
	config.Limit = newConfig.Limit
}
