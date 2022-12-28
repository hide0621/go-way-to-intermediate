package main

import (
	"database/sql"
	"fmt"

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

	// トランザクションを張る（この中でクエリを実行する）
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	article_id := 1
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`

	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback() //クエリのどこかで失敗したらロールバックして元に戻す
		return
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback() //クエリのどこかで失敗したらロールバックして元に戻す
		return
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback() //クエリのどこかで失敗したらロールバックして元に戻す
		return
	}

	tx.Commit() //クエリ全てが成功したのでその結果を反映させる
}
