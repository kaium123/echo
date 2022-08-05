package errors

import (
	"errors"
	"net/http"
)

var (
	ErrDataNotFound = errors.New("Data Not Found")
	ErrExist        = errors.New("Data already exist")
)

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

func ErrConflict(msg string) *ErrRes {
	return &ErrRes{
		Message: msg,
		Status:  http.StatusConflict,
		Error:   "That product already exist",
	}
}

func CheckErr(err error, msg string) *ErrRes {
	if err == ErrDataNotFound {
		return ErrNotFound(msg + " Not Found")
	}

	if err == ErrExist {
		return ErrConflict("Insert a Unique " + msg)
	}

	return ErrInternalServerErr("Some thing went wrong")
}
