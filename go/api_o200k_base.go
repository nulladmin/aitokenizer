// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nulladmin.com aitokenizer
 *
 * Aitokenizer is a Nulladmin.com microservice which performs input tokenization for AI workloads. It supports various encodings used by OpenAI, Anthropic and others.
 *
 * API version: 1.0.0
 */

package aitokenizer

import (
	"net/http"
	"os"
	"strings"
)

// O200kBaseAPIController binds http requests to an api service and writes the service results to the http response
type O200kBaseAPIController struct {
	service      O200kBaseAPIServicer
	errorHandler ErrorHandler
}

// O200kBaseAPIOption for how the controller is set up.
type O200kBaseAPIOption func(*O200kBaseAPIController)

// WithO200kBaseAPIErrorHandler inject ErrorHandler into controller
func WithO200kBaseAPIErrorHandler(h ErrorHandler) O200kBaseAPIOption {
	return func(c *O200kBaseAPIController) {
		c.errorHandler = h
	}
}

// NewO200kBaseAPIController creates a default api controller
func NewO200kBaseAPIController(s O200kBaseAPIServicer, opts ...O200kBaseAPIOption) *O200kBaseAPIController {
	controller := &O200kBaseAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the O200kBaseAPIController
func (c *O200kBaseAPIController) Routes() Routes {
	return Routes{
		"EncodeO200kBasePost": Route{
			strings.ToUpper("Post"),
			"/v1/encode/o200k_base",
			c.EncodeO200kBasePost,
		},
	}
}

// EncodeO200kBasePost - Upload text file for BPF encoding
func (c *O200kBaseAPIController) EncodeO200kBasePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var filenameParam *os.File
	{
		param, err := ReadFormFileToTempFile(r, "filename")
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "filename", Err: err}, nil)
			return
		}

		filenameParam = param
	}

	result, err := c.service.EncodeO200kBasePost(r.Context(), filenameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
