package errdefs

type ParentError interface {
	Status() int
}

type ErrNotFound interface {
	NotFoundError()
}

type ErrInvalidParameter interface {
	InvalidParameterError()
}

type ErrUnauthorized interface {
	UnauthorizedError()
}

type ErrForbidden interface {
	ForbiddenError()
}

type ErrSystem interface {
	SystemError()
}

type ErrInvalidData interface {
	InvalidDataError()
}

type ErrHttpStatus interface {
	StatusError()
}
