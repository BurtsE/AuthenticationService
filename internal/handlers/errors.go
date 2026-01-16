package handlers

import "errors"

var (
	ErrParsingRequest = errors.New("Error parsing request")
	ErrInternalServer = errors.New("Error processing request")
)
