package pagination_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func newCtx(path string) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

func toJson(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}
