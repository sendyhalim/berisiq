package helpers

import (
	"net/http"
)

func ServerError(w http.ResponseWriter, err error) {
	// TODO: Return json
	http.Error(w, err.Error(), 500)
}
