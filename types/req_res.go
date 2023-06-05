package types

type Response struct {
	Result        any            `json:"result"`
	ResponseStaus ResponseStatus `json:"responseStatus"`
}

type ResponseStatus struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}
