/*
@Time : 2019/9/20 10:13
@Author : mp
@File : sql
@Software: GoLand
*/
package main

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/common/log"
)

type SqlConfig struct {
	DriverName string `default:"mysql"`
	DataSource string `default:"root:123456@/service"`
}

type RtcService struct {
}

var (
	config *SqlConfig
)

var (
	m_db *sql.DB
)

func startDB()  {
	if err:=initSql();err!=nil{
		log.Fatal("init mysql failed！\n")
		return
	}
}

func initSql() error {
	if m_db != nil {
		return nil
	}

	db, err := sql.Open(config.DriverName, config.DataSource)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return err
	}

	m_db = db
	return nil
}

func doSqlExec(stmt *sql.Stmt, args ...interface{}) error {
	defer stmt.Close()
	res, err := stmt.Exec(args...)
	if err != nil {
		fmt.Printf("sql doExec :%s\n", err.Error())
		return err
	}
	// 通过LastInsertId可以获取插入数据的id
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Printf("sql doExec :%s\n", err.Error())
		return err
	}

	// 通过RowsAffected可以获取受影响的行数
	_, err = res.RowsAffected()
	if err != nil {
		fmt.Printf("sql doExec :%s\n", err.Error())
		return err
	}
	return nil
}

