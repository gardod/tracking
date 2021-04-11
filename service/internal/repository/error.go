package repository

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrNoResults = Error("no results found")
	ErrUnknown   = Error("unknown error occured")
)
