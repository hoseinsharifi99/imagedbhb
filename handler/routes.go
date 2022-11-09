package handler

import (
	"github.com/labstack/echo"
)

type Handler struct {
	ech *echo.Echo
}

func Newhandler() *Handler {
	h := &Handler{ech: echo.New()}
	h.defineRoute()
	return h
}

func (h *Handler) defineRoute() {

	h.ech.POST("/upload", Upload)
	h.ech.GET("/image/:gameID", GetImage)

}

func (h *Handler) Start() {
	h.ech.Logger.Fatal(h.ech.Start(":8000"))
}
