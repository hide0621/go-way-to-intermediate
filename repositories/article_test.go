package repositories_test

import (
	"database/sql"
	"fmt"
	"go-way-to-intermediate/models"
	"go-way-to-intermediate/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

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
			got, err := repositories.SelectArticleDetail(db, test.expected.ID)
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
