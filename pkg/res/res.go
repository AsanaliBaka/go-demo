package res

import (
	"encoding/json"
	"net/http"
)

func JsonWriter(w http.ResponseWriter, massage any, codeStatus int) {
	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(codeStatus)
	json.NewEncoder(w).Encode(massage)
}
