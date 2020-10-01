package errdefs

import (
	"errors"
	"fmt"
)

// not found error
type notFoundError struct {
	Code int
	Err  error
}

func (nf *notFoundError) NotFoundError() {}

func (nf *notFoundError) Error() string {
	return fmt.Sprintf("error: %v, code: %d", nf.Err, nf.Code)
}

func (ne *notFoundError) Status() int {
	return ne.Code
}

func NotFoundError(err error) error {
	if err == nil {
		return nil
	}

	return &notFoundError{Code: 404, Err: err}
}

// invalid data error
type invalidDataError struct {
	Code int
	Err  error
}

func (nf *invalidDataError) InvalidDataError() {}

func (nf *invalidDataError) Error() string {
	return fmt.Sprintf("error: %v, code: %d", nf.Err, nf.Code)
}

func (ine *invalidDataError) Status() int {
	return ine.Code
}

func InvalidDataError() error {
	return &invalidDataError{Code: 405, Err: errors.New("InvalidData")}
}

// http status error
type statusError struct {
	Code int
	Err  error
}

func (se *statusError) Error() string {
	return se.Err.Error()
}

func (se *statusError) Status() int {
	return se.Code
}

func (se *statusError) StatusError() {}

func NewStatusError(code int, err error) error {
	if err == nil {
		return nil
	}

	return &statusError{Code: code, Err: err}
}
