package action

import (
	"github.com/labstack/echo/v4"
	"github.com/ytake/draft/payload"
	"net/http"
)

// Ping
func (h *Handle) Ping(c echo.Context) error {
	// swagger:operation GET /ping ping getPing
	// For Server Status
	// ---
	//  produces:
	//  - application/json
	//  responses:
	//    200:
	//     $ref: "#/responses/PingResponse"
	return c.JSON(http.StatusOK, &payload.Ping{
		Status: "ok",
	})
}
