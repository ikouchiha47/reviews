package errdefs

func getImplementer(err error) error {
	switch err.(type) {
	case
		ParentError,
		ErrNotFound,
		ErrInvalidData,
		ErrInvalidParameter,
		ErrUnauthorized,
		ErrForbidden,
		ErrSystem:
		return err
	default:
		return err
	}
}

func IsHttpError(err error) bool {
	_, ok := getImplementer(err).(ErrHttpStatus)

	return ok
}

func IsNotFoundError(err error) bool {
	_, ok := getImplementer(err).(ErrNotFound)

	return ok
}

func IsInvalidDataError(err error) bool {
	_, ok := getImplementer(err).(ErrInvalidData)

	return ok
}

func GetCode(err error) int {
	_, ok := getImplementer(err).(ParentError)
	if ok {
		errr := err.(ParentError)
		return errr.Status()
	}

	return 0
}
