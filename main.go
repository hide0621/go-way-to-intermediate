package main

import (
	"go-way-to-intermediate/handlers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//別パッケージで定義したものを使うには「モジュール（プロジェクト）名/対象のディレクトリ(パッケージ)名」

func main() {

	// 明示的にルーターを使うことを宣言
	r := mux.NewRouter()

	//ここでHTTPメソッドと紐づけることが出来て、デフォルトで405エラーも返してくれる
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet) //パスパラメータで定義してハンドラ関数でそれを利用する
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r)) //第二引数にルーターを指定
}
