package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	ErrorCodeNotFound = "NOT_FOUND"
	ErrorInternal     = "INTERNAL"
)

func writeJson(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if v != nil {
		if err := json.NewEncoder(w).Encode(v); err != nil {
			panic(err)
		}
	}
}

func NotFound(w http.ResponseWriter) {
	writeJson(w, http.StatusNotFound, ErrorResponse{
		ErrorCode:    ErrorCodeNotFound,
		ErrorMessage: "Resource not found",
	})
}

func NoContent(w http.ResponseWriter) {
	writeJson(w, http.StatusNoContent, nil)
}

func Forbidden(w http.ResponseWriter, code string, err error) {
	writeJson(w, http.StatusForbidden, ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: err.Error(),
	})
}

func BadRequest(w http.ResponseWriter, code string, err error) {
	writeJson(w, http.StatusBadRequest, ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: err.Error(),
	})
}

func BadRequestWithDetails(w http.ResponseWriter, code string, err error, details string) {
	writeJson(w, http.StatusBadRequest, ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: err.Error(),
		Details:      details,
	})
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Print(err)

	writeJson(w, http.StatusInternalServerError, ErrorResponse{
		ErrorCode:    ErrorInternal,
		ErrorMessage: err.Error(),
	})
}

func Created(w http.ResponseWriter, v any) {
	writeJson(w, http.StatusCreated, v)
}

func RawOk(w http.ResponseWriter, v any) {
	writeJson(w, http.StatusOK, v)
}

func Ok(w http.ResponseWriter, v any, att string) {
	writeJson(w, http.StatusOK, map[string]any{att: v})
}
