package middlewares

import (
	"context"
	"log"
	"net/http"
)

// 自作のResponseWriterを作る
type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLogginWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeaderのオーバーライド
func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := req.Context()
		ctx = context.WithValue(ctx, traceIDKey{}, traceID)
		req = req.WithContext(ctx)
		rlw := NewResLogginWriter(w)

		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}
