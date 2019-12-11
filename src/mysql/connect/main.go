package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct{
	id int
	name string
	age int
}

func initDb() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_learn"
	db, err = sql.Open("mysql",dsn)
	//检测dsn格式是否✔，不会真正的校验数据库是否链接。
	if err != nil {
		return
	}
	err = db.Ping()
	//校验数据库链接是否成功
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
	return nil
}

func query(id int) (err error) {
	sqlStr := "select id,name,age from user where id = ?"
	var u user
	err = db.QueryRow(sqlStr,id).Scan(&u.id,&u.name,&u.age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
	return err
}

func queryMany() (err error) {
	sqlStr := "select * from user"
	rows, err := db.Query(sqlStr)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id,&u.name,&u.age)
		if err != nil {
			fmt.Printf("Scan Error:%v\n",err)
			return
		}
		fmt.Println(u)
	}
	return err

}

func insertOne() {
	sqlStr := "insert into user (name,age) values (?,?)"
	ret, err := db.Exec(sqlStr,"json",666)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func updateOne() {
	sqlStr := "update user set name = ? where id = ?"
	ret, err := db.Exec(sqlStr,"BOB",1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func deleteOne() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr,5)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)

	//回来
	//回来更新错误处理
	//回来更新遍历目录
	//回来更新mime类型
		//关于for循环包裹return
	//关于error向上暴露
}


	/**
		如何解决Sql注入，不要自己拼接字符串，使用预处理的方式，数据库的引擎会帮你做参数校验，不然就需要自己去写参数校验。
	 */
func PrepareInsert() {
	sqlStr := "insert into user (name,age) values (?,?)"
	stat, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("insert preparely failed, err:%v\n", err)
		return
	}
	defer stat.Close()
	_, err = stat.Exec("jack",100)
	_, err = stat.Exec("weston",78)
	if err != nil {
		fmt.Printf("insert preparely failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

func PrepareQueryMany() {
	sqlStr := "select * from user where id > ?"
	stat, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("query preparely failed, err:%v\n", err)
		return
	}
	rows,err := stat.Query(1)
	if err != nil {
		fmt.Printf("query preparely failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id,&u.name,&u.age)
		if err != nil {
			fmt.Printf("query preparely failed, err:%v\n", err)
			return
		}
		fmt.Println(u)
	}
}

func transcationDemo() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr := "update user set age = 11111 where id = ?"
	stat, err := tx.Prepare(sqlStr)
	if err != nil {
		tx.Rollback()
		fmt.Printf("Prepare failed, err:%v\n", err)
		return
	}
	_,err = stat.Exec(6)
	if err != nil {
		tx.Rollback()
		fmt.Printf("Prepare Exec failed, err:%v\n", err)
		return
	}
	sqlStr = "update user set w = 7777 where id = ?"
	_, err = tx.Exec(sqlStr,100)
	if err != nil {
		tx.Rollback()
		fmt.Printf("tx Exec failed, err:%v\n", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("Exec trans success")
	PrepareQueryMany()
}

func main() {
	fmt.Println([]rune("📕"))
	err := initDb()
	if err != nil {
		fmt.Println(err)
	}
	err = query(1)
	if err != nil {
		fmt.Println(err)
	}
	err = queryMany()
	if err != nil {
		fmt.Println(err)
	}
	insertOne()
	updateOne()
	deleteOne()
	PrepareInsert()
	PrepareQueryMany()
	transcationDemo()
}
