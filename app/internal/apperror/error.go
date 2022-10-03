package apperror

import "encoding/json"

var (
	ErrNotFound = NewAppError(nil, "not found", "", "US-000001")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message"`
	DeveloperService string `json:"developer_message"`
	Code             string `json:"code"`
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperService: developerMessage,
		Code:             code,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "system error", err.Error(), "US-000000")
}
