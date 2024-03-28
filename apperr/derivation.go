package apperr

type ClientError struct{ error }

func (e *ClientError) Unwrap() error {
	return e.error
}

func NewClientError(e error) *ClientError {
	return &ClientError{e}
}

type InternalError struct{ error }

func (e *InternalError) Unwrap() error {
	return e.error
}

func NewInternalError(e error) *InternalError {
	return &InternalError{e}
}

type NotFoundError struct{ error }

func (e *NotFoundError) Unwrap() error {
	return e.error
}

func NewNotFoundError(e error) *NotFoundError {
	return &NotFoundError{e}
}

type CanceledError struct{ error }

func NewCanceledError(e error) *CanceledError {
	return &CanceledError{e}
}

type ForbiddenError struct{ error }

func NewForbiddenError(e error) *ForbiddenError {
	return &ForbiddenError{e}
}

type UnauthorizedError struct{ error }

func NewUnauthorizedError(e error) *UnauthorizedError {
	return &UnauthorizedError{e}
}

type DeadlineExceededError struct{ error }

func NewDeadlineExceededError(e error) *DeadlineExceededError {
	return &DeadlineExceededError{e}
}

type RateLimitExceededError struct{ error }

func NewRateLimitExceededError(e error) *RateLimitExceededError {
	return &RateLimitExceededError{e}
}

type PreconditionFailedError struct{ error }

func NewPreconditionFailedError(e error) *PreconditionFailedError {
	return &PreconditionFailedError{e}
}
