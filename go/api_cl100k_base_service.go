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
	"context"
	"github.com/hupe1980/go-tiktoken"
	"net/http"
	"os"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.uber.org/zap"
)

// Cl100kBaseAPIService is a service that implements the logic for the Cl100kBaseAPIServicer
// This service should implement the business logic for every endpoint for the Cl100kBaseAPI API.
// Include any external packages or services that will be required by this service.
type Cl100kBaseAPIService struct {
}

// NewCl100kBaseAPIService creates a default api service
func NewCl100kBaseAPIService() *Cl100kBaseAPIService {
	return &Cl100kBaseAPIService{}
}

// EncodeCl100kBasePost - Upload text file for cl100k_base encoding
func (s *Cl100kBaseAPIService) EncodeCl100kBasePost(ctx context.Context, filename *os.File) (ImplResponse, error) {

	// It may be useful to log the temp file name passed in for debug purposes
	if viper.GetBool("debug") {
		zap.L().Debug("Encoding o200k_base", zap.String("filename", filename.Name()))
	}

	encoding, err := tiktoken.NewEncodingForModel("gpt-3.5")
	if err != nil {
		zap.L().Error("NewEncodingForMode error", zap.Error(err))
		r := &EncodeO200kBasePost400Response{Message: err.Error()}
		return Response(http.StatusInternalServerError, r), nil
	}

	fileContent, err := os.ReadFile(filename.Name())
	if err != nil {
		zap.L().Error("ReadFile error", zap.Error(err))
		r := &EncodeO200kBasePost400Response{Message: err.Error()}
		return Response(http.StatusInternalServerError, r), nil
	}

	// We no longer need the temp file so let's remove it.
	// Errors mean something is seriously wrong with the server
	// so we return an internal server error. Aruguably we could
	// just log, panic, and quit.
	err = os.Remove(filename.Name())
	if err != nil {
		zap.L().Error("Remove temp file error", zap.Error(err))
		r := &EncodeO200kBasePost400Response{Message: err.Error()}
		return Response(http.StatusInternalServerError, r), nil
	}

	ids, _, err := encoding.Encode(string(fileContent), nil, nil)
	if err != nil {
		zap.L().Error("Encode error", zap.Error(err))
		r := &EncodeO200kBasePost400Response{Message: err.Error()}
		return Response(http.StatusInternalServerError, r), nil
	}

	zap.L().Info("",
		zap.Int("Response", http.StatusOK),
	)

	return Response(http.StatusOK, ids), nil

}
