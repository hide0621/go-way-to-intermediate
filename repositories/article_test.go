package repositories_test

import (
	"go-way-to-intermediate/models"
	"go-way-to-intermediate/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {

	// テスト対象の関数を実行
	expectedNum := 3
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}
	// SelectArticleList関数から得たArticleスライスの長さが期待通りでないならFAILにする
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {

	// テーブルドリブンテストの準備
	tests := []struct {
		testTitle string         // テストのタイトル
		expected  models.Article // テストで期待する値
	}{
		{
			// 記事 ID1 番のテストデータ
			testTitle: "subtest1",
			expected: models.Article{
				ID:        1,
				Title:     "firstPost",
				Contetnts: "This is my first blog",
				UserName:  "saki",
				NiceNum:   3,
			},
		}, {
			// 記事 ID2 番のテストデータ
			testTitle: "subtest2",
			expected: models.Article{
				ID:        2,
				Title:     "2nd",
				Contetnts: "Second blog post",
				UserName:  "saki",
				NiceNum:   4,
			},
		},
	}

	// サブテストを一つずつまわす
	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contetnts != test.expected.Contetnts {
				t.Errorf("Content: get %s but want %s\n", got.Contetnts, test.expected.Contetnts)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}

		})
	}

}

func TestInsertArticle(t *testing.T) {

	article := models.Article{
		Title:     "insertTest",
		Contetnts: "testest",
		UserName:  "saki",
	}

	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum,
			newArticle.ID)
	}

	// テストでテーブルに追加したデータを削除しておく
	t.Cleanup(func() {
		const sqlstr = `
			delete from articles where title = ? and contents = ? and username = ? 
		`
		testDB.Exec(sqlstr, article.Title, article.Contetnts, article.UserName)
	})
}
