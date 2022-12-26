package main

import (
	"database/sql"
	"fmt"
	"go-way-to-intermediate/models"

	_ "github.com/go-sql-driver/mysql"
)

// import (
// 	"go-way-to-intermediate/handlers" //別パッケージで定義したものを使うには「モジュール（プロジェクト）名/対象のディレクトリ(パッケージ)名」
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {

// 	// 明示的にルーターを使うことを宣言
// 	r := mux.NewRouter()

// 	//ここでHTTPメソッドと紐づけることが出来て、デフォルトで405エラーも返してくれる
// 	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
// 	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
// 	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet) //パスパラメータで定義してハンドラ関数でそれを利用する
// 	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
// 	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

// 	log.Println("server start at port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", r)) //第二引数にルーターを指定
// }

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// データを挿入する処理
	article := models.Article{
		Title:     "insert test",
		Contetnts: "Can I insert data correctly?", // ここは自分で一から書いてあげないと「不定義」のエラーになる...
		UserName:  "saki",
	}
	const sqlStr = `
        insert into articles (title, contents, username, nice, created_at) values
        (?, ?, ?, 0, now());
    `

	result, err := db.Exec(sqlStr, article.Title, article.Contetnts, article.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 結果を確認
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

}
