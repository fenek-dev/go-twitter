package handlers

import "errors"

var (
	ErrCanNotGetUser = errors.New("could not get user from context")
)
