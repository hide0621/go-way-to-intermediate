package routers

//別パッケージで定義したものを使うには「モジュール（プロジェクト）名/対象のディレクトリ(パッケージ)名」
import (
	"go-way-to-intermediate/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(con *controllers.MyAppController) *mux.Router {

	// 明示的にルーターを使うことを宣言
	r := mux.NewRouter()

	//ここでHTTPメソッドと紐づけることが出来て、デフォルトで405エラーも返してくれる
	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet) //パスパラメータで定義してハンドラ関数でそれを利用する
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	return r

}
