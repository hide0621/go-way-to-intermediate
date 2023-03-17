package middlewares

import "sync"

var (
	logNo int = 1
	mu    sync.Mutex
)

// リクエストに対するトレースIDを生成する
func newTraceID() int {

	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}
