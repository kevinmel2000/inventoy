package errors

import (
	"context"
	"net/http"
)

func InternalServer(ctx context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"Status" : "Internal Server error"}`))
}

func DataNotFound(ctx context.Context, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{ "Status" : "Data not found!" }`))
}
