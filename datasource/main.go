package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	// sql
	var conStr string = "user=echo9z password=wzh19961217 dbname=ace_blog host=localhost sslmode=disable"

	// 打开数据库连接
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}

func QueryRes(db *sql.DB, query string) {
	// 查询数据
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	// 确保在函数返回时，无论正常返回还是出现错误，都会关闭 rows
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// 检查是否有错误
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
