package apperr

const (
	_ errorCode = userErrorCode

	UserNotFound errorCode = 12001
)

func userErrorText(code errorCode) string {
	switch code {
	case UserNotFound:
		return `ユーザーが見つかりません`
	default:
		return unknownErrorText
	}
}
