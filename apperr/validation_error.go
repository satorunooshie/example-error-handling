package apperr

const (
	_              errorCode = validationErrorCode
	InvalidRequest errorCode = 11001
)

func validationText(code errorCode) string {
	switch code {
	case InvalidRequest:
		return `不正なリクエストです`
	default:
		return unknownErrorText
	}
}
