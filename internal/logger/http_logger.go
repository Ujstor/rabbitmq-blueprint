package logger

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	Log.Errorf("Server error occurred: %v\n%s", err, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	Log.Errorf("Client error occurred: %s", http.StatusText(status))
	http.Error(w, http.StatusText(status), status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}