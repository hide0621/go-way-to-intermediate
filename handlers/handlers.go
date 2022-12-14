package handlers //mainパッケージ以外はファイルがあるディレクトリ名をパッケージ名として使う

import (
	"encoding/json"
	"go-way-to-intermediate/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)

}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// クエリパラメータpageを取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	// 暫定でこれを追加することで
	// 「変数pageが使われていない」というコンパイルエラーを回避
	log.Println(page)

	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

// ハンドラとパスの紐付け部分で定義したパスパラメータをハンドラ関数内で利用する
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 暫定でこれを追加することで
	// 「変数articleIDが使われていない」というコンパイルエラーを回避
	log.Println(articleID)

	article := models.Article1
	json.NewEncoder(w).Encode(article)

}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
