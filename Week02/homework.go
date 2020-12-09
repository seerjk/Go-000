package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	HttpRoute()
}

// API层：Http
func HttpRoute() {
	// 从URL Pramas获取id，假设id是1101
	id := 1101
	result, err := Service(id)
	if err != nil {
		// 错误时，接口返回固定内容，打印模拟
		fmt.Println("Welcome, nobody.")

		// 打印错误日志信息
		logrus.Warnf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		logrus.Warnf("stack trace: \n%+v\n", err)
	}

	// Http返回欢迎信息，这里用打印到屏幕模拟
	fmt.Println(result)
}

// Service层：生成欢迎信息
func Service(id int) (string, error) {
	name, err := QueryUsername(id)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Welcome, %s!", name)
	return result, nil
}

// Dao层：根据id查询name
func QueryUsername(id int) (string, error) {
	dbPath := "./foo.db"
	db, errOpen := sql.Open("sqlite3", dbPath)
	if errOpen != nil {
		return "", errors.Wrapf(errOpen, "failed to open %s", dbPath)
	}

	sqlStr := fmt.Sprintf("SELECT name FROM user WHERE id=%d Limit 1", id)
	rows, errQuery := db.Query(sqlStr)
	if errQuery == sql.ErrNoRows {
		return "", errors.Wrapf(errQuery, "failed to run sql: %s, no rows", sqlStr)
	} else if errQuery != nil {
		return "", errors.Wrapf(errQuery, "failed to run sql: %s", sqlStr)
	}

	var userName string
	errScan := rows.Scan(&userName)
	if errScan != nil {
		return "", errors.Wrapf(errScan, "failed to scan username")
	}

	return userName, nil
}
