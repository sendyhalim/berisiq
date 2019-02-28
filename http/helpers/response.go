package helpers

import (
	"net/http"

	"github.com/unrolled/render"
)

type jsonObj map[string]interface{}

var r = render.New()

func ServerError(w http.ResponseWriter, err error) {
	r.JSON(w, http.StatusInternalServerError, jsonObj{
		"error": jsonObj{
			"type":    "InternalServerError",
			"message": err.Error(),
		},
	})
}

func Success(w http.ResponseWriter, data interface{}) {
	r.JSON(w, http.StatusOK, jsonObj{
		"data": data,
	})
}
