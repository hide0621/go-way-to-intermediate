package api

//別パッケージで定義したものを使うには「モジュール（プロジェクト）名/対象のディレクトリ(パッケージ)名」
import (
	"database/sql"
	"go-way-to-intermediate/api/middlewares"
	"go-way-to-intermediate/controllers"
	"go-way-to-intermediate/services"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {

	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	// 明示的にルーターを使うことを宣言
	r := mux.NewRouter()

	// ここでHTTPメソッドと紐づけることが出来て、デフォルトで405エラーも返してくれる
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet) //パスパラメータで定義してハンドラ関数でそれを利用する
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)

	return r

}
