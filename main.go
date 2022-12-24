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

	articleID := 1
	const sqlStr = `
		select * 
		from articles
		where article_id = ? ; 
	`

	// rows, err := db.Query(sqlStr, articleID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer rows.Close()

	rows := db.QueryRow(sqlStr, articleID)
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var article models.Article
	var createdTime sql.NullTime
	// err := rows.Scan(&article.ID, &article.Title, &article.Contetnts, &article.UserName, &article.NiceNum, &article.CreatedAt)
	err = rows.Scan(&article.ID, &article.Title, &article.Contetnts, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)
}
