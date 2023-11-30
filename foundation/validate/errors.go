package validate

import "errors"

// error types should be concreted in here!!!

type ErrorResponse struct {
	Message string
}

// TODO: Unwrap Errors and handle them in errors middleware

func Unwrap(err error) error {
	var root error

	for err != nil {
		root = errors.Unwrap(err)
	}

	return root
}
