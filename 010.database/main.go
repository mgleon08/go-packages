package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// db 設定
const (
	userName = "root"
	password = ""
	host     = "127.0.0.1"
	port     = "3306"
	dbName   = "moai"
)

func initDB() {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// "username:password@tcp(host:port)/數據庫?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println(path)

	// 第一個是 driverName 第二個則是 database 的設定 path
	// 也可以用 var DB *sql.DB
	DB, _ := sql.Open("mysql", path)

	// 設定 database 最大連接數
	DB.SetConnMaxLifetime(100)

	//設定上 database 最大閒置連接數
	DB.SetMaxIdleConns(10)

	// 驗證是否連上 db
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail:", err)
		return
	}
	fmt.Println("connnect success")

	width := 640
	stmt, err := DB.Prepare("SELECT id FROM banner_specs WHERE width = ?")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(width)
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %v\n", id)
	}
}

func main() {
	initDB()
}
