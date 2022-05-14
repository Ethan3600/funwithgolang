package dtos

type R map[string]interface{}

type ResponseStatus string

const (
	Success ResponseStatus = "success"
	Fail    ResponseStatus = "fail"
	Error   ResponseStatus = "error"
)
