package util

import (
	"errors"
	"go-walet/domain"
)

func GetHttpStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	default:
		return 500
	}
}
