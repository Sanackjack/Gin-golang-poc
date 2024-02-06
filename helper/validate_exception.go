package helper

// CustomError is a custom error type
type ValidateException struct {
	Message string
	Code    int
}

// Error method to implement the error interface
func (e ValidateException) Error() string {
	return e.Message
}
