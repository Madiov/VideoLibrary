package ResponseHandler

type Response struct {
	StatusCode int
	Message    string
	Errors     string
	PayLoad    interface{}
}

func HandleResponse(err error) Response {
	return Response{}
}
