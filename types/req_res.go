package types

type Response struct {
	Result        any            `json:"result"`
	ResponseStaus ResponseStatus `json:"responseStatus"`
}

type ResponseStatus struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

func SuccessResponse(result any) (int, Response) {
	response := Response{
		Result: result,
		ResponseStaus: ResponseStatus{
			ErrorCode: 0,
			Message:   "",
		},
	}

	return 200, response
}

func ErrorResponse(errorMessage string, statusCode int) (int, Response) {
	response := Response{
		Result: nil,
		ResponseStaus: ResponseStatus{
			ErrorCode: 1,
			Message:   errorMessage,
		},
	}

	return statusCode, response
}
