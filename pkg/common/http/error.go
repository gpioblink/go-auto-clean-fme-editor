package http

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

// 使わないなら、安全に_にした方が良さそう
func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInternal(err error) render.Renderer {
	return &ErrResponse{
		Err:            nil,
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorText:      err.Error(),
	}
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            nil,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorText:      err.Error(),
	}
}
