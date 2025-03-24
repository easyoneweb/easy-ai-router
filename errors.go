package main

type HandlerErrors struct {
	OpenrouterErrors OpenrouterErrors
}

type OpenrouterErrors struct {
	Key      string
	ChatBody string
	Chat     string
}

var handlerErrors = HandlerErrors{
	OpenrouterErrors: OpenrouterErrors{
		Key:      "couldn't get key info",
		ChatBody: "check sent data",
		Chat:     "couldn't get chat data",
	},
}
