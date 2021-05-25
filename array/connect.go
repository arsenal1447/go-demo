package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	age  int
	name string
}

// 定义一个全局对象db
var db *sql.DB

func initDB() (err error) {
	dsn := "root:zxx123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// insertRowDemo()
	// queryRowDemo()
	// queryMutiRowDemo()
	updateRowDemo()
	fmt.Println("******************************\n")
	queryMutiRowDemo()
	fmt.Println("******************************\n")
	deleteRowDemo()
	queryMutiRowDemo()

}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select * from user where id=?"
	var u user
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name :%s age : %d\n", u.id, u.name, u.age)
}

func queryMutiRowDemo() {
	sqlStr := "select * from user where id >?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name :%s age : %d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user (name,age) values(?, ?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert ID failed, err:%v\n", err)
		return
	}

	fmt.Printf("insert success ,the id is %d\n", theId)
}

//  更新数据
func updateRowDemo() {
	sqlStr := "update user set age =? where id= ?"
	ret, err := db.Exec(sqlStr, 41, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}

	fmt.Printf("update success ,affected rows:%d\n", n)
}

// 插入数据
func deleteRowDemo() {
	sqlStr := "delete from user where id= ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}

	fmt.Printf("delete success ,affected rows:%d\n", n)
}
