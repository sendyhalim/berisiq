package notifications

import (
	"net/http"

	httpHelpers "github.com/sendyhalim/berisiq/http/helpers"
	"github.com/sendyhalim/berisiq/services"
)

func SendNotification(w http.ResponseWriter, r *http.Request) {
	err := services.Queue.SendMessage(
		"your queue url here~",
		"just for testing~",
	)

	if err != nil {
		httpHelpers.ServerError(w, err)

		return
	}

	w.Write([]byte("test"))
}
