package exception

type notFoundException struct {
	Code    int
	Message string
	Data    any
}

func NewNotFoundError(err error) notFoundException {
	response := notFoundException{
		Code:    404,
		Message: "NOT FOUND",
		Data: map[string]interface{}{
			"message": err.Error(),
		},
	}
	return response
}
