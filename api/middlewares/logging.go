package middlewares

import (
	"log"
	"net/http"
)

// 自作 ResponseWriter を作る
type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

// コンストラクタを作る
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// オーバーライドしたWriteHeaderメソッドを作る
func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

// ミドルウェアの中身
func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		traceID := newTraceID()

		// リクエスト情報をロギング
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		// 自作の ResponseWriter を作って
		rlw := NewResLoggingWriter(w)

		// それをハンドラに渡す
		next.ServeHTTP(rlw, req)

		// レスポンス情報をロギング
		log.Printf("[%d]res: %d", traceID, rlw.code)

	})
}
