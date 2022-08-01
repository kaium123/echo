package errors

import "net/http"

type ErrRes struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func ErrBadRequest(msg string) ErrRes {
	var err ErrRes
	err.Message = msg
	err.Status = http.StatusBadRequest
	err.Error = "Bad Request"
	return err
}

func ErrInternalServerErr(msg string) *ErrRes {

	return &ErrRes{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}

func ErrNotFound(msg string) *ErrRes {
	return &ErrRes{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "Content not found",
	}
}
