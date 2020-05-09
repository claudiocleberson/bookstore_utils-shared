package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/claudiocleberson/bookstore_utils-shared/utils/rest_err"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err rest_err.RestErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code())
	json.NewEncoder(w).Encode(err.ToJson())
}
