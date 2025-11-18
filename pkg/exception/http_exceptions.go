package exception


type AppError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Err   any    `json:"error,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, status int, err any) *AppError {
	return &AppError{
		Message: message,
		Status:  status,
		Err:   err,
	}
}
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Status:  404,
	}
}

func NewBadRequestError(err any) *AppError {
	return &AppError{
		Message: "Bad Request",
		Status:  400,
		Err:   err,
	}
}

func NewInternalServerError(message string, err any) *AppError {
	return &AppError{
		Message: message,
		Status:  500,
		Err:   err,
	}
}
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Message: message,
		Status:  401,
	}
}
func NewForbiddenError(message string) *AppError {
	return &AppError{
		Message: message,
		Status:  403,
	}
}
