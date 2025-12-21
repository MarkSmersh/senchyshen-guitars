package models

type ApiError struct {
	code    int
	content string
}

func newError(code int, content string) ApiError {
	return ApiError{
		code:    code,
		content: content,
	}
}

func (e ApiError) Code() int {
	return e.code
}

func (e ApiError) Error() string {
	return e.content
}

func (e ApiError) ErrorBytes() []byte {
	return []byte(e.content)
}

func (e ApiError) ToError() error {
	return e
}

func NewApiError(code int, content string) ApiError {
	return newError(code, content)
}

func BadRequest() ApiError {
	return NewApiError(400, "Bad Request")
}

func Unauthorized() ApiError {
	return NewApiError(401, "Unauthorized")
}

func NotFound() ApiError {
	return NewApiError(404, "Not Found")
}

func InternalServerError() ApiError {
	return NewApiError(500, "Internal Server Error")
}
