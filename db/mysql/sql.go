package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//1.0 exec
func insertUser(db *sql.DB) {
	_, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (?, ?)",
		"gopher",
		27,
	)
	check(err)
}

//2.0 Query
func query(db *sql.DB) {

	age := 27
	rows, err := db.Query("SELECT name FROM users WHERE age = ?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

//3.0 Prepared statements
func Prepared(db *sql.DB) {

	age := 27
	stmt, err := db.Prepare("SELECT name FROM users WHERE age = ?")
	check(err)
	rows, err := stmt.Query(age)
	defer rows.Close()
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}
	fmt.Println(record)
}

// 4.0  事务
/*
	    tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
*/
var db *sql.DB

/*
	其中连接参数可以有如下几种形式：
	user@unix(/path/to/socket)/dbname?charset=utf8
	user:password@tcp(localhost:5555)/dbname?charset=utf8
	user:password@/dbname
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	sql.Open函数实际上是返回一个连接池对象，不是单个连接。在open的时候并没有去连接数据库，只有在执行query、exce方法的时候才会去实际连接数据库。
	在一个应用中同样的库连接只需要保存一个sql.Open之后的db对象就可以了，不需要多次open。
	db 是一个*sql.DB类型的指针，在后面的操作中，都要用到db
	open之后，并没有与数据库建立实际的连接，与数据库建立实际的连接是通过Ping方法完成。
	此外，db应该在整个程序的生命周期中存在，也就是说，程序一启动，就通过Open获得db，直到程序结束，
	再Close db，而不是经常Open/Close。
*/
func connectDb() {
	//密码为空
	db, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	/*	连接池的实现关键在于SetMaxOpenConns和SetMaxIdleConns，其中：
		SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
		SetMaxIdleConns用于设置闲置的连接数。
	*/
	//	db.SetMaxOpenConns(2000)
	//	db.SetMaxIdleConns(1000)
	db.Ping()
}
func main() {
	db, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	defer db.Close()
	//insertUser(db)
	query(db)
	Prepared(db)
}
