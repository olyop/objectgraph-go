package objectgraph

import (
	"encoding/json"
	"net/http"
)

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var request GraphQLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the request
	response := e.Exec(r.Context(), &request)
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Other
	w.Header().Set("Content-Type", determineContentType(response))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func determineContentType(result *GraphQLResponse) string {
	if result.HasErrors() {
		return "application/problem+json; charset=utf-8"
	}

	return "application/json; charset=utf-8"
}
