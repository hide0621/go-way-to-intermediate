package apperrors

type MyAppError struct {
	// ErrCode型のErrCodeフィールド
	// (フィールド名を省略した場合、型名がそのままフィールド名になる)
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err     error  // エラーチェーンのための内部エラー
}

// Errorインターフェースを満たすように実装
func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// エラーチェーンにて、Unwrapメソッドを用いるための実装
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

// DB処理系のエラーをラップする
func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
