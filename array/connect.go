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
	// updateRowDemo()
	// fmt.Println("******************************\n")
	// queryMutiRowDemo()
	// fmt.Println("******************************\n")
	// deleteRowDemo()
	// queryMutiRowDemo()
	// prepareInsertDemo()

	// sqlInjectDemo("xxx' or 1=1#")
	// sqlInjectDemo("xxx' union select * from user #")
	// sqlInjectDemo("xxx' and (select count(*) from user) <10 #")

	transactionDemo()

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

// 查询多条数据示例
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

// 预处理查询示例 查询操作的预处理示例代码如下：
func prepareQueryDemo() {
	sqlStr := "select * from user where id>?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
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
		fmt.Printf("id :%d age :%d, name :%d \n", u.id, u.age, u.name)
	}
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name,age) values(?, ?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("boy", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("girl", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	fmt.Printf("insert success")
}

// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("selcet * from user where name ='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)

}

// 事务操作示例
func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "update user set age =22 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "update user set age =44 where id=?"
	ret2, err := tx.Exec(sqlStr2, 4)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	// fmt.Println("affRow1=="affRow1, affRow2)
	fmt.Printf("\n exec affRow1 : %d ", affRow1)
	fmt.Printf("\n exec affRow2 : %d \n", affRow2)

	if affRow1 == 1 && affRow2 == 1 {
		fmt.Printf("事务提交了")
		tx.Commit()
	} else {
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", sqlStr1)
		fmt.Printf("exec sql2 failed, err:%v\n", sqlStr2)
		fmt.Printf("事务回滚了")
	}
	fmt.Printf("exec trans success")
}
