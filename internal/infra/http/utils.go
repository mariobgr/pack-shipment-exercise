package http

import (
	"fmt"
	"net/http"

	"github.com/mariobgr/pack-shipment-exercise/internal/infra/logger"

	"github.com/go-chi/render"
)

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	// user-level status message
	StatusText string `json:"status" example:"Bad Request"`
	// application-level error message, for debugging
	ErrorText string `json:"error,omitempty" example:"something wrong"`
	// user-friendly description of the error
	Message string `json:"message,omitempty" example:"You do not have the necessary permissions to perform this action."`
	// grouped list of errors
	Errors map[string][]string `json:"errors,omitempty"`
}

// Render implements Renderer interface from go-chi/render
func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// SuccessResponse renderer type for handling success response.
type SuccessResponse struct {
	HTTPStatusCode int `json:"-"` // http response status code

	Success bool        `json:"success" example:"true"`
	Message interface{} `json:"message,omitempty"`
}

// Render implements Renderer interface from go-chi/render
func (e *SuccessResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func badRequest(writer http.ResponseWriter, request *http.Request, logger *logger.Logger, err error) {
	_ = render.Render(writer, request, ErrInvalidRequest(err))

	logger.Println("ERROR", "bad request", err)
}

// ErrInvalidRequest is a helper to create 400 error response
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, &ErrResponse{
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     fmt.Sprintf("not found"),
	})
}

func forbiddenResponse(writer http.ResponseWriter, request *http.Request) {
	_ = render.Render(writer, request, ErrForbidden)
}

// ErrForbidden is a 401 error response
var ErrForbidden = &ErrResponse{
	HTTPStatusCode: http.StatusUnauthorized,
	StatusText:     http.StatusText(http.StatusUnauthorized),
	Message:        "No API Key found in request",
}

func successResponse(writer http.ResponseWriter, request *http.Request, responseCode int, resource interface{}) {
	_ = render.Render(writer, request, &SuccessResponse{HTTPStatusCode: responseCode, Success: true, Message: resource})
}
