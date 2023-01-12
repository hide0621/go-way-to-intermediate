package repositories_test

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// テスト全体で共有する sql.DB 型
var testDB *sql.DB

// 全テスト共通の前処理を書く
func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// 前テスト共通の後処理を書く
func teardown() {
	testDB.Close()
}
