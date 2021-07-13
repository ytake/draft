package action

import (
	"github.com/labstack/echo/v4"
	"github.com/ytake/draft/payload"
	"github.com/ytake/draft/record"
	"net/http"
)

type (
	// Handle action / request handler
	Handle struct {
		record.Documenter
	}
	// Error action for default error response
	Error struct{}
)

type (
	Parameter struct {
		Key      record.DocumentKey     `json:"key"`
		Document map[string]interface{} `json:"document"`
		Expired  int32                  `json:"expired"`
	}
)

func (h *Handle) AddDocument(ctx echo.Context) error {
	p := new(Parameter)
	if err := ctx.Bind(p); err != nil {
		pe := &payload.ErrorResponse{Code: http.StatusInternalServerError}
		return ctx.JSON(http.StatusInternalServerError, pe.VngErrorFactory(
			err.Error(),
			payload.VndRequestURI(ctx.Request().RequestURI),
			ctx.Request().RequestURI,
		))
	}
	if err := h.Documenter.SaveDocument(record.WriteDocument{
		Document: record.Document{
			Key:  p.Key,
			Data: p.Document,
			Expire: p.Expired,
		},
	}); err != nil {
		pe := &payload.ErrorResponse{Code: http.StatusInternalServerError}
		return ctx.JSON(http.StatusInternalServerError, pe.VngErrorFactory(
			err.Error(),
			payload.VndRequestURI(ctx.Request().RequestURI),
			ctx.Request().RequestURI,
		))
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handle) FindDocument(ctx echo.Context) error {
	key := ctx.QueryParam("key")
	if key == "" {
		pe := &payload.ErrorResponse{Code: http.StatusBadRequest}
		return ctx.JSON(http.StatusInternalServerError, pe.VngErrorFactory(
			"bad request",
			payload.VndRequestURI(ctx.Request().RequestURI),
			ctx.Request().RequestURI,
		))
	}
	r, err := h.RetrieveDocument(record.DocumentKey(key))
	if err != nil {
		pe := &payload.ErrorResponse{Code: http.StatusInternalServerError}
		return ctx.JSON(http.StatusInternalServerError, pe.VngErrorFactory(
			err.Error(),
			payload.VndRequestURI(ctx.Request().RequestURI),
			ctx.Request().RequestURI,
		))
	}
	return ctx.JSON(http.StatusOK, r)
}
