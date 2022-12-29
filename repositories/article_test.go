package repositories_test

import (
	"database/sql"
	"fmt"
	"go-way-to-intermediate/models"
	"go-way-to-intermediate/repositories"
	"testing"
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

	expected := models.Article{
		ID:        1,
		Title:     "firstPost",
		Contetnts: "This is my first blog",
		UserName:  "saki",
		NiceNum:   2,
	}

	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
	}
	if got.Contetnts != expected.Contetnts {
		t.Errorf("Content: get %s but want %s\n", got.Contetnts, expected.Contetnts)
	}
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
}
