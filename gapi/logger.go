package gapi

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo, 
	handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		startTime := time.Now()
		result, err := handler(ctx, req)
		duration := time.Since(startTime)
		statusCode := codes.Unknown
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}

		logger := log.Info()
		if err != nil {
			logger = log.Error().Err(err)
		}

		logger.Str("protocal", "grpc").
			Str("method", info.FullMethod).
			Int("status_code", int(statusCode)).
			Str("status_text", statusCode.String()).
			Dur("duration", duration).
			Msg("reveived a gRPC request")
		return result, err
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body []byte
}

func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}

func HttpLogger(handlers http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		rec := &ResponseRecorder {
			ResponseWriter: res,
			StatusCode: http.StatusOK,
		}
		handlers.ServeHTTP(rec, req)
		duration := time.Since(startTime)

		logger := log.Info()
		if rec.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", rec.Body)
		}

		logger.Str("protocal", "http").
			Str("method", req.Method).
			Str("path", req.RequestURI).
			Int("status_code", rec.StatusCode).
			Str("status_text", http.StatusText(rec.StatusCode)).
			Dur("duration", duration).
			Msg("reveived a HTTP request")
	}) 
}