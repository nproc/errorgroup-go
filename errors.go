package errorgroup

import "strings"

// New creates a new ErrorGroup with the given errors
func New(errs []error) *ErrorGroup {
	return &ErrorGroup{
		Errors: errs,
	}
}

// ErrorGroup is used to group several errors together
type ErrorGroup struct {
	Errors []error
}

// Error returns all the error messages joined by a new line character
func (e *ErrorGroup) Error() string {
	strs := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		strs[i] = err.Error()
	}
	return strings.Join(strs, "\n")
}
