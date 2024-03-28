package apperr

const (
	unknownErrorText = `エラーが発生しました`

	group = 1000

	unknownErrorCode    = 10000
	validationErrorCode = 11000
	userErrorCode       = 12000
)

func text(code errorCode) string {
	switch group * (int(code) / group) {
	case validationErrorCode:
		return validationText(code)
	case userErrorCode:
		return userErrorText(code)
	}
	return unknownErrorText
}

const (
	UnknownError errorCode = unknownErrorCode
)
