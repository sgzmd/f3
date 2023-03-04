package main

import "fmt"

type ErrorCodeType int

const (
	NoUserId    ErrorCodeType = iota
	NoEntryId   ErrorCodeType = iota
	NoEntryType ErrorCodeType = iota
	NoEntities  ErrorCodeType = iota
)

type RequestError struct {
	error
	ErrorCode ErrorCodeType
	Message   string
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("RequestError: ErrorCode=%d, message=%s", r.ErrorCode, r.Message)
}

var ErrorCodeToMessage = map[ErrorCodeType]string{
	NoUserId:    "No UserId",
	NoEntryId:   "No EntryId",
	NoEntryType: "No EntryType",
	NoEntities:  "No Entities",
}

func createRequestError(code ErrorCodeType) *RequestError {
	return &RequestError{ErrorCode: code, Message: ErrorCodeToMessage[code]}
}
