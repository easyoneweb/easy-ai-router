package main

type HandlerErrors struct {
	OpenrouterErrors OpenrouterErrors
}

type OpenrouterErrors struct {
	Key      string
	Limits   string
	ChatBody string
	Chat     string
}

var handlerErrors = HandlerErrors{
	OpenrouterErrors: OpenrouterErrors{
		Key:      "couldn't get key info",
		Limits:   "couldn't get limits data",
		ChatBody: "check sent data",
		Chat:     "couldn't get chat data",
	},
}
