package service

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrAccountValidation = Error("invalid or inactive account ID")
	ErrAccountRetrieve   = Error("failed to retrieve account")
	ErrMessagePublish    = Error("failed to publish message")
)
