package utils

import (
	ErrorCode "travel/config"
)

func GetMessageFromErrorCode(errorCode int) string {
	message := ""
	switch errorCode {
	case ErrorCode.USER_EXITS:
		message = ErrorCode.USER_EXITS_MESSAGE
		break
	}

	return message
}
