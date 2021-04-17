package action

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/ytake/draft/payload"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandle_Ping(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/ping", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handle{}
	if h.Ping(c) != nil {
		t.Error("action handler error")
	}
	var p payload.Ping
	err := json.Unmarshal(rec.Body.Bytes(), &p)
	if err != nil {
		t.Error("it is not the expected response")
	}
	if p.Status != "ok" {
		t.Error("it is not the expected payload")
	}
}
