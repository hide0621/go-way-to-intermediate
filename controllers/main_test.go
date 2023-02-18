package controllers_test

import (
	"go-way-to-intermediate/controllers"
	"go-way-to-intermediate/controllers/testdata"
	"testing"
)

// テストに使うリソース (コントローラ構造体) を用意
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {

	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
