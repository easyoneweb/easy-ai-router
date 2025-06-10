package openrouter

type OpenrouterErrors struct {
	CreateRequest string
	DoRequest     string
	StatusCode    string
	ReadBody      string
	UnmarshalJson string
	MarshalJson   string
	LimitLog      string
}

var openrouterErrors = OpenrouterErrors{
	CreateRequest: "couldn't create request",
	DoRequest:     "request failed",
	StatusCode:    "code must be 200",
	ReadBody:      "couldn't read response body",
	UnmarshalJson: "couldn't unmarshal json body",
	MarshalJson:   "couldn't marshal post body",
	LimitLog:      "couldn't create limit log",
}
