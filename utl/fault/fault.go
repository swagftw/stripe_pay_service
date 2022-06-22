package fault

type HTTPError struct {
	Status  int
	ErrCode string
	Message string
	Res     string
	Err     error
	Service string
}

func New(statusCode int, service, msg, res, errCode string, err error) error {
	return &HTTPError{
		Status:  statusCode,
		Message: msg,
		Res:     res,
		Err:     err,
		Service: service,
	}
}

type ErrResponse struct {
	StatusCode int         `json:"status"`
	Message    interface{} `json:"message"`
	Res        string      `json:"res"`
	Service    string      `json:"service,omitempty"`
	Err        string      `json:"error"`
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}
