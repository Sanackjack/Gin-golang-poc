package helper

// CustomError is a custom error type
type CustomException struct {
	Message string
	Code    int
}

// Error method to implement the error interface
func (e CustomException) Error() string {
	return e.Message
}
