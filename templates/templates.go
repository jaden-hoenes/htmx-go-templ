package templates

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type TemplateHandler[T any] struct {
	Method  string
	Path    string
	Handler func(context echo.Context, binding T) (templ.Component, error)
}

func (h *TemplateHandler[T]) EchoHandler() (echo.HandlerFunc, error) {
	return func(c echo.Context) error {
		var binding T
		err := c.Bind(&binding)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		component, err := h.Handler(c, binding)
		if err != nil {
			return err
		}

		return component.Render(context.Background(), c.Response().Writer)
	}, nil
}

func (handler *TemplateHandler[T]) Register(e *echo.Echo) {
	h, _ := handler.EchoHandler()

	switch {
	case handler.Method == http.MethodGet:
		e.GET(handler.Path, h)
	case handler.Method == http.MethodPut:
		e.PUT(handler.Path, h)
	case handler.Method == http.MethodPost:
		e.POST(handler.Path, h)
	case handler.Method == http.MethodDelete:
		e.DELETE(handler.Path, h)
	case handler.Method == http.MethodPatch:
		e.PATCH(handler.Path, h)
	case handler.Method == http.MethodHead:
		e.HEAD(handler.Path, h)
	case handler.Method == http.MethodOptions:
		e.OPTIONS(handler.Path, h)
	case handler.Method == http.MethodConnect:
		e.CONNECT(handler.Path, h)
	case handler.Method == http.MethodTrace:
		e.TRACE(handler.Path, h)
	default:
	}
}
