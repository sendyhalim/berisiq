package notifications

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sendyhalim/berisiq/services"
)

func SendNotification(c echo.Context) error {
	err := services.Queue.SendMessage(
		"your queue url here~",
		"just for testing~",
	)

	if err != nil {
		c.Logger().Error(err)
	}

	return c.String(http.StatusOK, "done")
}
