package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Result - сообщение о результате
type Result struct {
	Msg interface{} `json:"result"`
}

// Error - сообщение об ошибке
type Error struct {
	Msg string `json:"error"`
}

//http ответ
func jsonResponse(err bool, w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if err {
		errResp := Error{msg.(string)}
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			http.Error(w, "responseJson error", http.StatusInternalServerError)
		}
	} else {
		resResp := Result{msg}
		if err := json.NewEncoder(w).Encode(resResp); err != nil {
			http.Error(w, "responseJson error", http.StatusInternalServerError)
		}
	}
}

// ValidateQuery для валидации метода и URL-query
func ValidateQuery(w http.ResponseWriter, r *http.Request, validateQuery ...string) bool {
	if r.Method != validateQuery[0] {
		jsonResponse(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("bad %v method", r.Method))
		return false
	}
	for _, v := range validateQuery[1:] {
		if !r.URL.Query().Has(v) {
			jsonResponse(true, w, http.StatusServiceUnavailable, "not parameters "+v)
			return false
		}
	}
	return true
}
